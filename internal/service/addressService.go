package service

import (
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

type AddressService struct {
}

// CreateAddress 创建地址信息
func (*AddressService) CreateAddress(ctx *gin.Context, request dto.AddressRequest) res.Response {
	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	address := &model.Address{
		UserID:  user.ID,
		Name:    request.Name,
		Phone:   request.Phone,
		Address: request.Address,
	}

	var addressDao dao.AddressDao
	err := addressDao.CreateAddress(address)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "地址创建失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "地址创建成功",
	}
}

// GetAddress 查看地址信息
func (*AddressService) GetAddress(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	addressID, err := strconv.Atoi(id)
	if err != nil {
		util.Loggers.Errorln(err)
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
	var addressDao dao.AddressDao
	address, err := addressDao.GetAddressByID(uint(addressID), user.ID)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "地址查询失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: vo.BuildAddressResponse(address),
		Msg:  "地址查询成功",
	}
}

// ListAddress 查看所有地址
func (*AddressService) ListAddress(ctx *gin.Context) res.Response {
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
	addresses, err := addressDao.ListAddressByUserID(user.ID)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "地址查询失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildAddressResponseList(addresses), int64(len(addresses))),
		Msg:  "地址查询成功",
	}
}

// UpdateAddress 修改地址信息
func (*AddressService) UpdateAddress(ctx *gin.Context, request dto.AddressRequest) res.Response {
	id := ctx.Param("id")
	addressID, err := strconv.Atoi(id)
	if err != nil {
		util.Loggers.Errorln(err)
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
	address := &model.Address{
		Name:    request.Name,
		Phone:   request.Phone,
		Address: request.Address,
	}
	var addressDao dao.AddressDao
	err = addressDao.UpdateAddressByID(uint(addressID), user.ID, address)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "地址修改失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "地址修改成功",
	}
}

// DeleteAddress 删除地址信息
func (*AddressService) DeleteAddress(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	addressID, err := strconv.Atoi(id)
	if err != nil {
		util.Loggers.Errorln(err)
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
	var addressDao dao.AddressDao
	err = addressDao.DeleteAddressByID(uint(addressID), user.ID)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "地址删除失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "地址删除成功",
	}
}
