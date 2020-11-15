package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/api/handler"
	"websocket/api/middleware"
	"websocket/api/response"
	"websocket/config"
)

type Router struct {
	middleware *middleware.MiddleWare
	handler    *handler.Handler
}

func NewRouter(middleware *middleware.MiddleWare, handler *handler.Handler) *Router {
	return &Router{middleware: middleware, handler: handler}
}

func (router *Router) Run() *gin.Engine {
	r := gin.Default()

	r.Use(router.middleware.Translations())
	// api
	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/detail", router.handler.Api.Detail) // 详情信息
		apiRouter.POST("/login", router.handler.Api.Login)  // 用户登录
	}

	// 用户相关
	userRouter := r.Group("/user", router.middleware.AccessToken())
	userRouter.POST("/ws", router.handler.User.Ws)

	// websocket 处理
	r.GET("/ws/:app_id", router.handler.Ws.WebSocket)

	pushRouter := r.Group("/ws/push", router.middleware.Sign())
	{
		pushRouter.POST("/user", router.handler.Push.User)  // 发送到指定用户
		pushRouter.POST("/many", router.handler.Push.User)  // 发送到多个用户
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
