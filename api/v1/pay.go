package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
)

// OrderPay 订单支付
func OrderPay(ctx *gin.Context) {
	var payRequest dto.PayRequest
	err := ctx.ShouldBind(&payRequest)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
	}
	var payService service.PayService
	response := payService.OrderPay(ctx, payRequest)
	ctx.JSON(http.StatusOK, response)
}
