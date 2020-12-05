// +build wireinject

package main

import (
	"github.com/google/wire"
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
	serviceApi "websocket/service/api"
)

func NewHttp(router2 *router.Router) *http.Server {
	return &http.Server{
		Addr:           config.App.Address,
		Handler:        router2.Run(),
		MaxHeaderBytes: 1 << 20,
	}
}

var providerSet = wire.NewSet(
	// connection
	connection.NewMySQL,
	connection.NewRedis,

	cache.NewUserCache,

	// repo
	repo.NewApp,
	repo.NewMessage,
	repo.NewUser,
	repo.NewMessageRead,

	// api handler
	api.NewApi,
	push.NewPush,
	user.NewUser,
	handler.NewWs,

	// api 服务
	serviceApi.NewMessageService,

	// middleWare
	middleware.NewMiddleWare,

	// router
	router.NewRouter,

	wire.Struct(new(handler.Handler), "*"),

	NewHttp,
)

func Initialize() *http.Server {
	panic(wire.Build(providerSet))
}
