package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/api/router"
	"websocket/config"
	"websocket/connection"
	"websocket/global"
	"websocket/service"
)

func main() {

	// 连接mysql数据库
	global.NewConnect("default")

	// 连接redis数据库
	connection.NewRedisDb("default")

	go service.GlobalHub.Run()

	// 开启调试模式
	if config.App.Debug == "on" {
		gin.SetMode("debug")
	}

	handler := router.NewRouter()

	s := &http.Server{
		Addr:           config.App.Address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Http Listen" + config.App.Address)
	s.ListenAndServe()
}
