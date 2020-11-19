package main

import (
	"fmt"
	"net/http"

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

	router := Initialize()
	fmt.Println("Http Listen" + config.App.Address)
	s := &http.Server{
		Addr:           config.App.Address,
		Handler:        router.Run(),
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
