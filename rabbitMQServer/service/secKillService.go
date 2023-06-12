package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"mall-demo/internal/initialize"
	"mall-demo/internal/model"
	"mall-demo/internal/pkg/util"
	"mall-demo/internal/rabbitMQ"

	"gorm.io/gorm"
)

func SecKillService() {
	// 获取通道
	ch, err := rabbitMQ.MQ.Channel()
	if err != nil {
		panic(err)
	}

	// 声明队列
	queue, err := ch.QueueDeclare("secKill_queue", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// 限速,每次一个
	err = ch.Qos(1, 0, false)
	if err != nil {
		panic(err)
	}

	// 消费
	msgs, err := ch.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	// 监听生产者，需要阻塞进程
	go func() {
		fmt.Println("开始监听....")
		for msg := range msgs {
			var secKillOrder model.SecKillOrder
			// 反序列化订单信息
			err = json.Unmarshal(msg.Body, &secKillOrder)
			if err != nil {
				panic(err)
			}
			opt := &sql.TxOptions{
				Isolation: sql.LevelReadCommitted, // 事务隔离级别
			}
			// 开启事务
			tx := initialize.DB.Begin(opt)
			// 减库存
			result := tx.Model(&model.SecKillProduct{}).
				Where("id = ? and num >= ?", secKillOrder.SecKillProductID, secKillOrder.Num).
				Update("num", gorm.Expr("num - ?", secKillOrder.Num))
			if result.Error != nil {
				tx.Rollback()
				util.Loggers.Errorln("减库存失败")
				return
			}
			if result.RowsAffected == 0 {
				tx.Rollback()
				util.Loggers.Errorln("减库存失败")
				return
			}
			// 创建订单
			err = tx.Model(&model.SecKillOrder{}).Create(&secKillOrder).Error
			if err != nil {
				tx.Rollback()
				util.Loggers.Errorln("创建订单失败")
				return
			}
			tx.Commit()
			util.Loggers.Infoln("创建订单成功")

			// 确认消息已消费
			_ = msg.Ack(false)
		}
	}()
}
