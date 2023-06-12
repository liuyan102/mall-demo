package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dao"
	"mall-demo/internal/dto"
	"mall-demo/internal/model"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/vo"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
}

func (*OrderService) CreateOrder(ctx *gin.Context, request dto.OrderRequest) res.Response {
	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	var addressDao dao.AddressDao
	address, err := addressDao.GetAddressByID(request.AddressID, user.ID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "地址不存在",
		}
	}

	number := fmt.Sprintf("%09v%d%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000),
		request.ProductID, request.UserID)
	orderNum, _ := strconv.ParseUint(number, 10, 64)

	order := &model.Order{
		UserID:    user.ID,
		BossID:    request.BossID,
		ProductID: request.ProductID,
		AddressID: address.ID,
		Num:       request.Num,
		OrderNum:  orderNum,
		Money:     request.Money,
		Type:      0, // 默认未支付
	}

	var orderDao dao.OrderDao
	err = orderDao.CreateOrder(order)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "订单创建失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "订单创建成功",
	}
}

func (*OrderService) ShowOrder(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		return res.Response{
			Code: e.InvalidParam,
			Data: nil,
			Msg:  "参数错误",
		}
	}

	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	var orderDao dao.OrderDao
	order, err := orderDao.GetOrderByID(uint(orderID), user.ID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "订单不存在",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: vo.BuildOrderResponse(&order),
		Msg:  "订单查询成功",
	}

}

func (*OrderService) ListOrder(ctx *gin.Context, request dto.OrderRequest) res.Response {
	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)
	if request.PageSize == 0 {
		request.PageSize = 15
	}

	var orderDao dao.OrderDao
	condition := make(map[string]interface{})
	if request.Type == 0 || request.Type == 1 {
		condition["type"] = request.Type
	}
	condition["user_id"] = user.ID

	orderList, err := orderDao.ListOrderByCondition(condition, request.BasePage)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "订单查询失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildOrderResponseList(orderList), int64(len(orderList))),
		Msg:  "订单查询成功",
	}
}

func (*OrderService) DeleteOrder(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	orderID, err := strconv.Atoi(id)
	if err != nil {
		return res.Response{
			Code: e.InvalidParam,
			Data: nil,
			Msg:  "参数错误",
		}
	}

	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	var orderDao dao.OrderDao
	err = orderDao.DeleteOrderByID(uint(orderID), user.ID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "订单删除失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "订单删除成功",
	}
}
