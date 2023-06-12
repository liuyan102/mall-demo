package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
)

// CreateFavorite 创建收藏夹
func CreateFavorite(ctx *gin.Context) {
	var favoriteService service.FavoriteService
	var createFavoriteRequest dto.CreateFavoriteRequest
	err := ctx.ShouldBind(&createFavoriteRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	response := favoriteService.CreateFavorite(ctx, createFavoriteRequest)
	ctx.JSON(http.StatusOK, response)
}

// ShowFavorite 展示收藏夹
func ShowFavorite(ctx *gin.Context) {
	var favoriteService service.FavoriteService
	response := favoriteService.ShowFavorite(ctx)
	ctx.JSON(http.StatusOK, response)
}

// DeleteFavorite 删除收藏夹
func DeleteFavorite(ctx *gin.Context) {
	var favoriteService service.FavoriteService
	response := favoriteService.DeleteFavorite(ctx)
	ctx.JSON(http.StatusOK, response)
}
