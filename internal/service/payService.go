package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dao"
	"mall-demo/internal/dto"
	"mall-demo/internal/model"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/pkg/util"
	"mall-demo/internal/vo"
	"strconv"
)

type PayService struct {
}

// OrderPay 订单支付
func (*PayService) OrderPay(ctx *gin.Context, request dto.PayRequest) res.Response {
	orderDao := dao.NewOrderDaoWithDB()
	// 开启事务
	tx := orderDao.Begin()

	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	order, err := orderDao.GetOrderByID(request.OrderID, user.ID)
	if err != nil {
		util.Loggers.Error(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "订单不存在",
		}
	}

	money := order.Money * float64(order.Num)

	// 对用户金额进行解密，减去订单，再加密保存
	userMoneyStr, err := util.AesDecrypt(user.Money, request.UserKey)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "用户密钥错误",
		}
	}
	userMoney, _ := strconv.ParseFloat(userMoneyStr, 64)
	if userMoney-money < 0.0 {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "余额不足",
		}
	}
	finMoney := fmt.Sprintf("%f", userMoney-money)
	user.Money, _ = util.AesEncrypt(finMoney, request.UserKey)

	var userDao dao.UserDao
	err = userDao.UpdateUserMoneyByID(user.ID, user.Money)
	if err != nil {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "支付失败",
		}
	}

	// 对商家金额进行解密，增加订单金额，再加密保存
	bossMoneyStr, err := util.AesDecrypt(order.Boss.Money, request.BossKey)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商家密钥错误",
		}
	}
	bossMoney, _ := strconv.ParseFloat(bossMoneyStr, 64)
	finMoney = fmt.Sprintf("%f", bossMoney+money)
	order.Boss.Money, _ = util.AesEncrypt(finMoney, request.BossKey)
	err = userDao.UpdateUserMoneyByID(order.Boss.ID, order.Boss.Money)
	if err != nil {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商家密钥错误",
		}
	}

	// 对应商品数量-1
	var productDao dao.ProductDao
	order.Product.Num -= order.Num
	err = productDao.UpdateProductNumByID(order.Product.ID, order.Product.Num)
	if err != nil {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品数量更新失败",
		}
	}
	order.Type = 1
	// 修改订单状态为已支付 type = 1
	err = orderDao.UpdateOrderTypeByID(user.ID, order.ID, order.Type)
	if err != nil {
		tx.Rollback()
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "订单状态更新失败",
		}
	}
	tx.Commit()
	return res.Response{
		Code: e.Success,
		Data: vo.BuildPayResponse(&order),
		Msg:  "支付成功",
	}
}
