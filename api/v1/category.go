package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/service"
	"net/http"
)

// ListCategory 展示商品分类
func ListCategory(ctx *gin.Context) {
	var categoryService service.CategoryService
	response := categoryService.ListCategory(ctx)
	ctx.JSON(http.StatusOK, response)
}
