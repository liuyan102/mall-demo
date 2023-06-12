package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
)

// CreateOrder 创建订单
func CreateOrder(ctx *gin.Context) {
	var createOrderRequest dto.OrderRequest
	err := ctx.ShouldBind(&createOrderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var orderService service.OrderService
	response := orderService.CreateOrder(ctx, createOrderRequest)
	ctx.JSON(http.StatusOK, response)
}

func ListOrder(ctx *gin.Context) {
	var listOrderRequest dto.OrderRequest
	err := ctx.ShouldBind(&listOrderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var orderService service.OrderService
	response := orderService.ListOrder(ctx, listOrderRequest)
	ctx.JSON(http.StatusOK, response)
}

func ShowOrder(ctx *gin.Context) {
	var orderService service.OrderService
	response := orderService.ShowOrder(ctx)
	ctx.JSON(http.StatusOK, response)
}

func DeleteOrder(ctx *gin.Context) {
	var orderService service.OrderService
	response := orderService.DeleteOrder(ctx)
	ctx.JSON(http.StatusOK, response)
}
