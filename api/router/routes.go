package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websocket/api/response"
	"websocket/config"
)

func (router *Router) Run() *gin.Engine {
	r := gin.Default()

	r.Use(router.middleware.Translations())

	// 后端接口
	adminRouter := r.Group("/admin")
	{
		// 用户信息
		adminRouter.POST("/user/list", router.handler.AdminUser.List)
		adminRouter.POST("/user/create", router.handler.AdminUser.Create)
		adminRouter.POST("/user/update", router.handler.AdminUser.Update)
		adminRouter.POST("/user/offline", router.handler.AdminUser.Offline)
		adminRouter.POST("/user/online", router.handler.AdminUser.Online)
		adminRouter.POST("/user/delete", router.handler.AdminUser.Delete)

		// 应用信息
		adminRouter.POST("/app/list", router.handler.AdminApp.List)
		adminRouter.POST("/app/create", router.handler.AdminApp.Create)
		adminRouter.POST("/app/update", router.handler.AdminApp.Update)
		adminRouter.POST("/app/offline", router.handler.AdminApp.Offline)
		adminRouter.POST("/app/online", router.handler.AdminApp.Online)
		adminRouter.POST("/app/delete", router.handler.AdminApp.Delete)
	}

	// api
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/detail", router.handler.Api.Detail)      // 详情信息
		apiRouter.POST("/login", router.handler.Api.Login)       // 用户登录
		apiRouter.POST("/register", router.handler.Api.Register) // 用户注册
	}

	// 用户相关
	userRouter := r.Group("/user", router.middleware.AccessToken())
	userRouter.POST("/ws", router.handler.User.Ws)

	// websocket 处理
	r.GET("/ws/:app_id", router.handler.Ws.WebSocket)

	pushRouter := r.Group("/ws/push", router.middleware.Sign())
	{
		pushRouter.POST("/user", router.handler.Push.User)  // 发送到指定用户
		pushRouter.POST("/many", router.handler.Push.Many)  // 发送到多个用户
		pushRouter.POST("/all", router.handler.Push.User)   // 发送到全部用户
		pushRouter.POST("/group", router.handler.Push.User) // 发送到指定分组
	}

	// 页面没有找到
	r.NoRoute(func(context *gin.Context) {
		response.BusinessError(context, "404 Not Found")
	})

	// 前端文件
	r.StaticFS("/public", http.Dir(config.App.StaticUrl))

	return r
}
