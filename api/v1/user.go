package v1

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dto"
	"mall-demo/internal/service"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var userService service.UserService
	var userRegisterReq dto.UserRegisterRequest
	err := ctx.ShouldBind(&userRegisterReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	response := userService.UserRegister(ctx, userRegisterReq)
	ctx.JSON(http.StatusOK, response)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var userService service.UserService
	var userLoginReq dto.UserLoginRequest
	err := ctx.ShouldBind(&userLoginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	response := userService.UserLogin(ctx, userLoginReq)
	ctx.JSON(http.StatusOK, response)
}

// UserUpdate 修改信息
func UserUpdate(ctx *gin.Context) {
	var userService service.UserService
	NickName := ctx.PostForm("NickName")
	response := userService.NickNameUpdate(ctx, NickName)
	ctx.JSON(http.StatusOK, response)
}

// AvatarUpload 头像上传
func AvatarUpload(ctx *gin.Context) {
	var userService service.UserService
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	fileSize := fileHeader.Size
	response := userService.AvatarUpload(ctx, file, fileSize)
	ctx.JSON(http.StatusOK, response)
}

// SendEmail 发送邮箱
func SendEmail(ctx *gin.Context) {
	var userService service.UserService
	var sendEmailRequest dto.SendEmailRequest
	err := ctx.ShouldBind(&sendEmailRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	response := userService.SendEmail(ctx, sendEmailRequest)
	ctx.JSON(http.StatusOK, response)
}

// ValidEmail 验证邮箱
func ValidEmail(ctx *gin.Context) {
	var userService service.UserService
	response := userService.ValidEmail(ctx)
	ctx.JSON(http.StatusOK, response)
}

// ShowMoney 显示金额
func ShowMoney(ctx *gin.Context) {
	var userService service.UserService
	key := ctx.PostForm("Key")
	response := userService.ShowMoney(ctx, key)
	ctx.JSON(http.StatusOK, response)
}
