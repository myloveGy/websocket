package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websocket/api/handler"
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
	webSocket := r.Group("/ws")
	{
		webSocket.GET("/:app_id", handler.WebSocket)
		webSocket.POST("/push/user", handler.MessagePush)  // 发送到指定用户
		webSocket.POST("/push/many", handler.MessagePush)  // 发送到多个用户
		webSocket.POST("/push/all", handler.MessagePush)   // 发送到全部用户
		webSocket.POST("/push/group", handler.MessagePush) // 发送到指定用户
	}

	// 前端文件
	r.StaticFS("/public", http.Dir(config.App.StaticUrl))

	return r
}
