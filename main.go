package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"websocket/config"
	"websocket/service"
)

func main() {

	go service.GlobalHub.Run()

	// 开启调试模式
	if config.App.Debug == "on" {
		gin.SetMode("debug")
	}

	s := Initialize()
	fmt.Println("Http Listen" + config.App.Address)
	s.ListenAndServe()
}
