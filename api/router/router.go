package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/api/handler/api"
	"websocket/api/handler/push"
	"websocket/api/handler/user"
	"websocket/cache"
	"websocket/connection"

	"websocket/api/handler"
	"websocket/api/middleware"
	"websocket/config"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	redisConnection := connection.NewRedis()
	userCache := cache.NewUserCache(redisConnection)
	middle := middleware.NewMiddleWare(userCache)
	r.Use(middle.Translations())
	// api
	apiRouter := r.Group("/api")
	{
		apiHandler := api.NewApi(userCache)
		apiRouter.GET("/detail", apiHandler.Detail) // 详情信息
		apiRouter.POST("/login", apiHandler.Login)  // 用户登录
	}

	// 用户相关
	userRouter := r.Use(middle.AccessToken())
	userHandler := user.NewUser()
	userRouter.POST("/user/ws", userHandler.Ws)

	// websocket 处理
	r.GET("/ws/:app_id", handler.WebSocket)

	pushRouter := r.Group("/ws/push")
	pushRouter.Use(middle.Sign())
	{
		pushHandler := &push.Push{}
		pushRouter.POST("/user", pushHandler.User)  // 发送到指定用户
		pushRouter.POST("/many", pushHandler.User)  // 发送到多个用户
		pushRouter.POST("/all", pushHandler.User)   // 发送到全部用户
		pushRouter.POST("/group", pushHandler.User) // 发送到指定分组
	}

	// 前端文件
	r.StaticFS("/public", http.Dir(config.App.StaticUrl))

	return r
}
