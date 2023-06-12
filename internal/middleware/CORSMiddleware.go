package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method               // 请求方法
		origin := ctx.Request.Header.Get("Origin") // 请求头

		if origin != "" {
			// 必填，接收指定域的请求可以使用*不加以限制，但是不安全，且不允许XMLHttpRequest携带Cookie
			// ctx.Writer.Header().Set("Access-Control-Allow-Origin","*")
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
			// 必填，设置服务器支持的所有跨域请求的方法
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
			// 如果浏览器请求包括Access-Control-Request-Headers字段，则Access-Control-Allow-Headers字段是必需的。
			// 它也是一个逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段。
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
			// 可选,XMLHttpRequest的响应对象能拿到的额外字段
			ctx.Writer.Header().Set("Access-Control-Expose-Header", "Access-Control-Allow-Headers")
			// 可选，是否允许后续请求携带认证信息Cookie，该值只能是true，不需要则不设置
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			// 返回响应数据格式
			ctx.Set("content-type", "application/json")
		}

		// 放行所有的options方法
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 处理请求
		ctx.Next()
	}
}
