package service

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dao"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/res"
	"mall-demo/internal/pkg/util"
	"mall-demo/internal/vo"
)

type CarouselService struct {
}

func (*CarouselService) ListCarousel(ctx *gin.Context) res.Response {
	var carouselDao dao.CarouselDao
	list, err := carouselDao.List()
	if err != nil {
		util.Loggers.Errorln(err)
		return res.Response{
			Code: e.Error,
			Data: nil,
			Msg:  "获取失败",
		}
	}
	util.Loggers.Infoln("listCarousel success")
	return res.Response{
		Code: e.Success,
		Data: res.BuildDataList(vo.BuildCarouselListResponse(list), int64(len(list))),
		Msg:  "获取成功",
	}
}
