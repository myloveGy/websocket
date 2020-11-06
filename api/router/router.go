package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websocket/api/handler"
	"websocket/api/middleware"
	"websocket/config"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// api
	api := r.Group("/api")
	{
		api.GET("/info", handler.ApiInfo)      // 详情信息
		api.POST("/login", handler.ApiLogin)   // 用户登录
		api.POST("/verify", handler.ApiVerify) // 验证登录
	}

	// websocket 处理
	r.GET("/ws/:app_id", handler.WebSocket)

	push := r.Group("/ws/push")
	push.Use(middleware.Sign())
	{
		push.POST("/user", handler.WsPushUser)  // 发送到指定用户
		push.POST("/many", handler.WsPushUser)  // 发送到多个用户
		push.POST("/all", handler.WsPushUser)   // 发送到全部用户
		push.POST("/group", handler.WsPushUser) // 发送到指定用户
	}

	// 前端文件
	r.StaticFS("/public", http.Dir(config.App.StaticUrl))

	return r
}
