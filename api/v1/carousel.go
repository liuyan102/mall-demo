package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/service"
	"net/http"
)

// ListCarousel 展示轮播图
func ListCarousel(ctx *gin.Context) {
	var carousel service.CarouselService
	response := carousel.ListCarousel(ctx)
	ctx.JSON(http.StatusOK, response)
}
