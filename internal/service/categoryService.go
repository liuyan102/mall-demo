package service

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dao"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/pkg/util"
	"mall-demo/internal/vo"
)

type CategoryService struct {
}

// ListCategory 展示商品分类
func (*CategoryService) ListCategory(ctx *gin.Context) res.Response {
	var categoryDao dao.CategoryDao
	categoryList, err := categoryDao.ListCategory()
	if err != nil {
		util.Loggers.Error(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "查询失败",
		}
	}
	util.Loggers.Info("list category success")
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildCategoryResponseList(categoryList), int64(len(categoryList))),
		Msg:  "查询成功",
	}
}
