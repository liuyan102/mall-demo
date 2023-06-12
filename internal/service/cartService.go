package service

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dao"
	"mall-demo/internal/dto"
	"mall-demo/internal/model"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/vo"
	"strconv"
)

type CartService struct {
}

func (*CartService) CreateCart(ctx *gin.Context, request dto.CartRequest) res.Response {
	var productDao dao.ProductDao
	_, err := productDao.GetProductByID(request.ProductID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品不存在",
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

	cart := &model.Cart{
		UserID:    user.ID,
		ProductID: request.ProductID,
		BossID:    request.BossID,
		Num:       request.Num,
		MaxNum:    100,
		Check:     false,
	}

	var cartDao dao.CartDao
	err = cartDao.CreateCart(cart)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "加入购物车失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "加入购物车成功",
	}

}

func (*CartService) ShowCart(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	cartID, err := strconv.Atoi(id)
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

	var cartDao dao.CartDao
	cart, err := cartDao.GetCartByID(uint(cartID), user.ID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品不存在",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: vo.BuildCartResponse(cart),
		Msg:  "查询成功",
	}
}

func (*CartService) ListCart(ctx *gin.Context) res.Response {
	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	var cartDao dao.CartDao
	cartList, err := cartDao.LisCartByUserID(user.ID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品不存在",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildCartResponseList(cartList), int64(len(cartList))),
		Msg:  "查询成功",
	}
}

func (*CartService) UpdateCart(ctx *gin.Context, request dto.CartRequest) res.Response {
	id := ctx.Param("id")
	cartID, err := strconv.Atoi(id)
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

	var cartDao dao.CartDao
	err = cartDao.UpdateCartNumByID(uint(cartID), user.ID, request.Num)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "修改失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "修改成功",
	}
}

func (*CartService) DeleteCart(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	cartID, err := strconv.Atoi(id)
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

	var cartDao dao.CartDao
	err = cartDao.DeleteCartByID(uint(cartID), user.ID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品删除失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "商品删除成功",
	}
}
