package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
)

func CreateCart(ctx *gin.Context) {
	var createCartRequest dto.CartRequest
	err := ctx.ShouldBind(&createCartRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var cartService service.CartService
	response := cartService.CreateCart(ctx, createCartRequest)
	ctx.JSON(http.StatusOK, response)
}

func ShowCart(ctx *gin.Context) {
	var cartService service.CartService
	response := cartService.ShowCart(ctx)
	ctx.JSON(http.StatusOK, response)
}

func ListCart(ctx *gin.Context) {
	var cartService service.CartService
	response := cartService.ListCart(ctx)
	ctx.JSON(http.StatusOK, response)
}

func UpdateCart(ctx *gin.Context) {
	var updateCartRequest dto.CartRequest
	err := ctx.ShouldBind(&updateCartRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var cartService service.CartService
	response := cartService.UpdateCart(ctx, updateCartRequest)
	ctx.JSON(http.StatusOK, response)
}

func DeleteCart(ctx *gin.Context) {
	var cartService service.CartService
	response := cartService.DeleteCart(ctx)
	ctx.JSON(http.StatusOK, response)
}
