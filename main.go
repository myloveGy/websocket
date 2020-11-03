package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/api/handler/ws"
	"websocket/api/router"
	"websocket/config"
	"websocket/global"
)

func main() {

	global.NewConnect("default")

	go ws.GlobalHub.Run()

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
