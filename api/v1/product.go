package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
)

// CreateProduct 创建商品
func CreateProduct(ctx *gin.Context) {
	var productService service.ProductService
	var createProductRequest dto.CreateProductRequest
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	files := form.File["file"]
	err = ctx.ShouldBind(&createProductRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	response := productService.CreateProduct(ctx, createProductRequest, files)
	ctx.JSON(http.StatusOK, response)

}

// ListProduct 展示商品
func ListProduct(ctx *gin.Context) {
	var productService service.ProductService
	var listProductRequest dto.ListProductRequest
	err := ctx.ShouldBind(&listProductRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	response := productService.ListProduct(ctx, listProductRequest)
	ctx.JSON(http.StatusOK, response)
}

// SearchProduct 搜索商品
func SearchProduct(ctx *gin.Context) {
	var productService service.ProductService
	var searchProductRequest dto.SearchProductRequest
	err := ctx.ShouldBind(&searchProductRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	response := productService.SearchProduct(ctx, searchProductRequest)
	ctx.JSON(http.StatusOK, response)
}

// ShowProduct 查看商品详细信息
func ShowProduct(ctx *gin.Context) {
	var productService service.ProductService
	response := productService.ShowProduct(ctx)
	ctx.JSON(http.StatusOK, response)
}

// ShowProductImg 查看商品图片
func ShowProductImg(ctx *gin.Context) {
	var productService service.ProductService
	response := productService.ShowProductImg(ctx)
	ctx.JSON(http.StatusOK, response)
}
