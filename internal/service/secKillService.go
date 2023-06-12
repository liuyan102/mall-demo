package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mall-demo/internal/cache"
	"mall-demo/internal/dao"
	"mall-demo/internal/dto"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/pkg/util"
	"mall-demo/internal/rabbitMQ"
	"mall-demo/internal/vo"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/redis/go-redis/v9"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type SecKillService struct {
}

// AddSecKillProduct 添加秒杀商品
func (*SecKillService) AddSecKillProduct(ctx *gin.Context, request dto.SecKillProductRequest, files []*multipart.FileHeader) res.Response {
	user, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	boss := user.(*model.User)
	// 以第一张图片作为封面图
	firstFile, err := files[0].Open()
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.InvalidParam,
			Data: nil,
			Msg:  "文件打开失败",
		}
	}
	secKillProductPath, err := util.UploadProduct(firstFile, boss.ID, request.Name)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品图片上传失败",
		}
	}
	secKillProduct := &model.SecKillProduct{
		Name:         request.Name,
		CategoryID:   request.CategoryID,
		Title:        request.Title,
		Info:         request.Info,
		ImgPath:      secKillProductPath,
		Price:        request.Price,
		SecKillPrice: request.SecKillPrice,
		Num:          request.Num,
		BossID:       boss.ID,
		BossName:     boss.UserName,
		BossAvatar:   boss.Avatar,
	}
	var secKillProductDao dao.SecKillProductDao
	err = secKillProductDao.CreateSecKillProduct(secKillProduct)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品创建失败",
		}
	}
	// 将要秒杀的商品库存存入redis
	key := fmt.Sprintf("secKillProduct_%v", secKillProduct.ID)
	if err = cache.RedisDB.Set(ctx.Request.Context(), key, secKillProduct.Num, 0).Err(); err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "存入缓存失败",
		}
	}

	util.Loggers.Infoln("create secKillProduct success")
	return res.Response{
		Code: e.Success,
		Data: vo.BuildSecKillProductResponse(secKillProduct),
		Msg:  "秒杀商品创建成功",
	}
}

// SecKillWithoutLock 无锁秒杀
func (*SecKillService) SecKillWithoutLock(ctx *gin.Context, request dto.SecKill) res.Response {
	opt := &sql.TxOptions{
		Isolation: sql.LevelReadCommitted, // 事务隔离级别为 serializable 时会发生死锁，原因是读写锁竞争导致的，每一个select获得读锁，每一个update获得写锁
	}
	tx := initialize.DB.Begin(opt)

	// 检查登录状态
	value, exists := ctx.Get("user")
	if !exists {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	// 检查库存
	var num int
	err := tx.
		Model(&model.SecKillProduct{}).
		Select("num").
		Where("id=?", request.SecKillProductID).
		First(&num).Error
	fmt.Println(num)

	if err != nil {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品不存在",
		}
	}
	if num == 0 {
		tx.Rollback()
		return res.Response{
			Code: e.Success,
			Data: nil,
			Msg:  "商品已售罄",
		}
	}
	if request.Num > num {
		tx.Rollback()
		return res.Response{
			Code: e.Success,
			Data: nil,
			Msg:  "商品库存不足",
		}
	}

	// 扣库存
	if num >= request.Num {
		result := tx.Model(&model.SecKillProduct{}).
			Where("id = ? and num >= ?", request.SecKillProductID, request.Num).
			Update("num", gorm.Expr("num - ?", request.Num))
		if result.Error != nil {
			tx.Rollback()
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "库存更新失败",
			}
		}
		if result.RowsAffected == 0 {
			tx.Rollback()
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "库存更新失败",
			}
		}

		// 创建订单
		number := fmt.Sprintf("%09v%d%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000),
			request.SecKillProductID, user.ID)
		secKillOrderNum, _ := strconv.ParseUint(number, 10, 64)

		secKillOrder := &model.SecKillOrder{
			UserID:           user.ID,
			SecKillProductID: request.SecKillProductID,
			BossID:           request.BossID,
			AddressID:        request.AddressID,
			Num:              request.Num,
			SecKillOrderNum:  secKillOrderNum,
			Type:             0,
			Money:            request.Money,
		}

		err = tx.Model(&model.SecKillOrder{}).Create(&secKillOrder).Error
		if err != nil {
			tx.Rollback()
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "创建订单失败",
			}
		}
	}
	tx.Commit()
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "秒杀成功",
	}

}

// SecKillWithXLock 排他锁秒杀
func (*SecKillService) SecKillWithXLock(ctx *gin.Context, request dto.SecKill) res.Response {
	opt := &sql.TxOptions{
		Isolation: sql.LevelReadCommitted, // 事务隔离级别
	}
	tx := initialize.DB.Begin(opt)
	// 检查登录状态
	value, exists := ctx.Get("user")
	if !exists {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)
	// 检查库存
	var num int
	err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Model(&model.SecKillProduct{}).
		Select("num").
		Where("id=?", request.SecKillProductID).
		First(&num).Error
	fmt.Println(num)
	if err != nil {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品不存在",
		}
	}
	if num == 0 {
		tx.Rollback()
		return res.Response{
			Code: e.Success,
			Data: nil,
			Msg:  "商品已售罄",
		}
	}
	if request.Num > num {
		tx.Rollback()
		return res.Response{
			Code: e.Success,
			Data: nil,
			Msg:  "商品库存不足",
		}
	}
	// 扣库存
	if num >= request.Num {
		result := tx.Model(&model.SecKillProduct{}).
			Where("id = ? and num >= ?", request.SecKillProductID, request.Num).
			Update("num", gorm.Expr("num - ?", request.Num))
		if result.Error != nil {
			tx.Rollback()
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "库存更新失败",
			}
		}
		if result.RowsAffected == 0 {
			tx.Rollback()
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "库存更新失败",
			}
		}
		// 创建订单
		number := fmt.Sprintf("%09v%d%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000),
			request.SecKillProductID, user.ID)
		secKillOrderNum, _ := strconv.ParseUint(number, 10, 64)

		secKillOrder := &model.SecKillOrder{
			UserID:           user.ID,
			SecKillProductID: request.SecKillProductID,
			BossID:           request.BossID,
			AddressID:        request.AddressID,
			Num:              request.Num,
			SecKillOrderNum:  secKillOrderNum,
			Type:             0,
			Money:            request.Money,
		}

		err = tx.Model(&model.SecKillOrder{}).Create(&secKillOrder).Error
		if err != nil {
			tx.Rollback()
			return res.Response{
				Code: e.Error,
				Data: nil,
				Msg:  "创建订单失败",
			}
		}
	}
	tx.Commit()
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "秒杀成功",
	}

}

// SecKillWithRedis redis秒杀
func (*SecKillService) SecKillWithRedis(ctx *gin.Context, request dto.SecKill, count int) res.Response {
	// 检查登录状态
	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	// 检查库存
	// 减库存,并将用户订单信息保存进redis
	// 采用lua脚本实现原子操作
	script := redis.NewScript(`
		local secKillProductID = KEYS[1]
		local num = tonumber(ARGV[1])
		local userID = ARGV[2]
		local count = ARGV[3]
		
		local stockKey = "secKillProduct_" .. secKillProductID
		local userKey = "secKillProduct_" .. secKillProductID .. "_users_" .. count
		
		local stock = tonumber(redis.call("get",stockKey))
		-- 检查库存
		if stock < num then
			return 0
		end

		-- 减库存
    	redis.call("decrby", stockKey, num)
    
    	-- 写入用户信息
    	redis.call("set", userKey, userID)
    	return stock
    `)
	result, err := script.Run(ctx.Request.Context(), cache.RedisDB,
		[]string{strconv.Itoa(int(request.SecKillProductID))}, request.Num, user.ID, count).Result()
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "秒杀失败",
		}
	}
	if result.(int64) == 0 {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "库存不足",
		}
	}
	fmt.Println(count, result)

	// 生产者将订单信息存入消息队列，等待消费者将信息写入mysql数据库

	// 创建通道
	ch, err := rabbitMQ.MQ.Channel()
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "创建通道失败",
		}
	}

	// 声明消息要发送的队列
	queue, err := ch.QueueDeclare(
		"secKill_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "创建队列失败",
		}
	}

	// 创建订单
	number := fmt.Sprintf("%09v%d%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000),
		request.SecKillProductID, user.ID)
	secKillOrderNum, _ := strconv.ParseUint(number, 10, 64)

	secKillOrder := &model.SecKillOrder{
		UserID:           user.ID,
		SecKillProductID: request.SecKillProductID,
		BossID:           request.BossID,
		AddressID:        request.AddressID,
		Num:              request.Num,
		SecKillOrderNum:  secKillOrderNum,
		Type:             0,
		Money:            request.Money,
	}

	// 序列化消息主体
	body, _ := json.Marshal(secKillOrder)

	// 将消息发布到声明的队列
	err = ch.PublishWithContext(ctx.Request.Context(), "", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "消息发送失败",
		}
	}
	// 生产者消息发布成功，等待

	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "秒杀成功",
	}
}
