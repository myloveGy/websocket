package main

import (
	"fmt"

	"websocket/config"
	"websocket/service"

	"github.com/gin-gonic/gin"
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
