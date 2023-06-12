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

type FavoriteService struct {
}

// CreateFavorite 创建收藏夹
func (*FavoriteService) CreateFavorite(ctx *gin.Context, request dto.CreateFavoriteRequest) res.Response {
	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	var productDao dao.ProductDao
	product, err := productDao.GetProductByID(request.ProductID)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品不存在",
		}
	}

	var bossDao dao.UserDao
	boss, err := bossDao.GetUserInfoByID(request.BossID)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商家不存在",
		}
	}

	var favoriteDao dao.FavoriteDao
	exist, _ := favoriteDao.Exist(user.ID, product.ID)
	if exist {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "商品已存在收藏夹",
		}
	}

	favorite := &model.Favorite{
		User:      *user,
		UserID:    user.ID,
		Product:   *product,
		ProductID: request.ProductID,
		Boss:      *boss,
		BossID:    request.BossID,
	}

	err = favoriteDao.Create(favorite)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "收藏失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "收藏成功",
	}
}

// ShowFavorite 展示收藏夹
func (*FavoriteService) ShowFavorite(ctx *gin.Context) res.Response {
	value, exists := ctx.Get("user")
	if !exists {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "登录过期",
		}
	}
	user := value.(*model.User)

	var favoriteDao dao.FavoriteDao
	favoriteList, err := favoriteDao.List(user.ID)
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "查询失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildFavoriteResponseList(favoriteList), int64(len(favoriteList))),
		Msg:  "查询成功",
	}
}

// DeleteFavorite 删除收藏夹
func (*FavoriteService) DeleteFavorite(ctx *gin.Context) res.Response {
	id := ctx.Param("id")
	favoriteID, err := strconv.Atoi(id)
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
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
	var favoriteDao dao.FavoriteDao
	err = favoriteDao.Delete(user.ID, uint(favoriteID))
	if err != nil {
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "删除失败",
		}
	}
	return res.Response{
		Code: e.Success,
		Data: nil,
		Msg:  "删除成功",
	}
}
