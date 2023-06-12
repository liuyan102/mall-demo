package middleware

import (
	"github.com/gin-gonic/gin"
	"mall-demo/internal/dao"
	"mall-demo/internal/pkg/e"
	"mall-demo/internal/pkg/util"
	"net/http"
	"strings"
)

// JWTAuthMiddleware 认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求头header中的Authorization
		token := ctx.GetHeader("Authorization")
		// 判断token format是不是Bearer格式
		if token == "" || !strings.HasPrefix(token, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Code": e.UnAuthorized,
				"Msg":  "权限不足",
			})
			// 抛弃请求
			ctx.Abort()
			return
		}

		// 取Bearer之后的部分进行验证
		token = token[7:]
		// 解析token
		claims, err := util.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Code": e.UnAuthorized,
				"Msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		// 验证通过
		userID := claims.ID
		// 查找userID是否存在
		var userDao dao.UserDao
		user, err := userDao.GetUserInfoByID(userID)
		// 用户不存在
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Code": e.UnAuthorized,
				"Msg":  "权限不足",
			})
		}

		// 用户存在，将用户信息写入上下文
		ctx.Set("user", user)

		// 继续执行后面的请求
		ctx.Next()
	}
}
