package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/api/handler"
	"websocket/api/handler/api"
	"websocket/api/handler/push"
	"websocket/api/handler/user"
	"websocket/api/middleware"
	"websocket/api/router"
	"websocket/cache"
	"websocket/config"
	"websocket/connection"
	"websocket/repo"
	"websocket/service"
)

func main() {

	go service.GlobalHub.Run()

	// 开启调试模式
	if config.App.Debug == "on" {
		gin.SetMode("debug")
	}

	mysqlConnection := connection.NewDB()
	redisConnection := connection.NewRedis()
	userCache := cache.NewUserCache(redisConnection)

	appRepo := repo.NewApp(mysqlConnection)
	userRepo := repo.NewUser(mysqlConnection)
	messageRepo := repo.NewMessage(mysqlConnection)
	messageReadRepo := repo.NewMessageRead(mysqlConnection)

	apiHandler := api.NewApi(userCache, userRepo)
	pushHandler := push.NewPush(messageRepo, messageReadRepo)
	userHandler := user.NewUser()
	wsHandler := handler.NewWs(appRepo)

	middle := middleware.NewMiddleWare(userCache, appRepo)
	handle := handler.NewHandler(apiHandler, pushHandler, userHandler, wsHandler)
	routerHandler := router.NewRouter(middle, handle)

	s := &http.Server{
		Addr:           config.App.Address,
		Handler:        routerHandler.Run(),
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Http Listen" + config.App.Address)
	s.ListenAndServe()
}
