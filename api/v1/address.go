package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
)

// CreateAddress 创建地址
func CreateAddress(ctx *gin.Context) {
	var createAddressRequest dto.AddressRequest
	err := ctx.ShouldBind(&createAddressRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var addressService service.AddressService
	response := addressService.CreateAddress(ctx, createAddressRequest)
	ctx.JSON(http.StatusOK, response)
}

// GetAddress 查看地址信息
func GetAddress(ctx *gin.Context) {
	var addressService service.AddressService
	response := addressService.GetAddress(ctx)
	ctx.JSON(http.StatusOK, response)
}

// ListAddress 展示所有地址
func ListAddress(ctx *gin.Context) {
	var addressService service.AddressService
	response := addressService.ListAddress(ctx)
	ctx.JSON(http.StatusOK, response)
}

// UpdateAddress 修改地址信息
func UpdateAddress(ctx *gin.Context) {
	var updateAddressRequest dto.AddressRequest
	err := ctx.ShouldBind(&updateAddressRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	var addressService service.AddressService
	response := addressService.UpdateAddress(ctx, updateAddressRequest)
	ctx.JSON(http.StatusOK, response)
}

// DeleteAddress 删除地址信息
func DeleteAddress(ctx *gin.Context) {
	var addressService service.AddressService
	response := addressService.DeleteAddress(ctx)
	ctx.JSON(http.StatusOK, response)
}
