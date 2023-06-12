package router

import (
	api "mall-demo/api/v1"
	"mall-demo/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.StaticFS("/static", http.Dir("./static")) // 加载静态资源
	v1 := r.Group("api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
		// 用户注册登录
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)

		// 用户模块操作
		authed := v1.Group("/") // 需要登录保护
		authed.Use(middleware.JWTAuthMiddleware())
		{
			// 修改用户信息
			authed.PUT("user", api.UserUpdate)
			// 上传头像
			authed.POST("avatar", api.AvatarUpload)
			// 发送邮件
			authed.POST("user/mail", api.SendEmail)
			// 验证邮件
			authed.POST("user/valid/mail", api.ValidEmail)
			// 显示金额
			authed.POST("money", api.ShowMoney)

			// 创建商品
			authed.POST("product/create", api.CreateProduct)
			// 搜索商品
			authed.POST("product/search", api.SearchProduct)

			// 展示收藏夹
			authed.GET("favorite", api.ShowFavorite)
			// 创建收藏夹
			authed.POST("favorite", api.CreateFavorite)
			// 删除收藏夹
			authed.DELETE("favorite/:id", api.DeleteFavorite)

			// 创建地址
			authed.POST("address", api.CreateAddress)
			// 查看地址
			authed.GET("address/:id", api.GetAddress)
			// 展示地址
			authed.GET("address", api.ListAddress)
			// 修改地址
			authed.PUT("address/:id", api.UpdateAddress)
			// 删除地址
			authed.DELETE("address/:id", api.DeleteAddress)

			// 加入购物车
			authed.POST("cart", api.CreateCart)
			// 查看购物车中商品信息
			authed.GET("cart/:id", api.ShowCart)
			// 展示购物车
			authed.GET("cart", api.ListCart)
			// 修改购物车信息
			authed.PUT("cart/:id", api.UpdateCart)
			// 删除购物车商品
			authed.DELETE("cart/:id", api.DeleteCart)

			// 创建订单
			authed.POST("order", api.CreateOrder)
			// 展示订单
			authed.GET("order", api.ListOrder)
			// 查看订单
			authed.GET("order/:id", api.ShowOrder)
			// 删除订单
			authed.DELETE("order/:id", api.DeleteOrder)

			// 支付订单
			authed.POST("pay", api.OrderPay)

			// 商品秒杀
			// 录入秒杀商品信息
			authed.POST("secKill/add", api.AddSecKillProduct)
			// 秒杀商品
			// 无锁 出现超卖
			authed.POST("secKill/withoutLock", api.SecKillWithoutLock)
			// 互斥锁秒杀 正常
			authed.POST("secKill/withMutexLock", api.SecKillWithMutexLock)
			// 悲观锁/排他锁秒杀 正常
			authed.POST("secKill/withXLock", api.SecKillWithXLock)
			// 悲观锁/排他锁秒杀 正常
			authed.POST("secKill/withRedis", api.SecKillWithRedis)

		}

		// 轮播图模块操作
		v1.GET("/carousel", api.ListCarousel)

		// 展示商品
		v1.GET("/product", api.ListProduct)
		// 查看商品详细信息
		v1.GET("/product/:id", api.ShowProduct)
		// 查看商品图片
		v1.GET("/img/:id", api.ShowProductImg)

		// 商品分类
		v1.GET("/category", api.ListCategory)
	}

	return r
}
