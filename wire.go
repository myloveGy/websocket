// +build wireinject

package main

import (
	"github.com/google/wire"

	"websocket/api/handler"
	"websocket/api/handler/api"
	"websocket/api/handler/push"
	"websocket/api/handler/user"
	"websocket/api/middleware"
	"websocket/api/router"
	"websocket/cache"
	"websocket/connection"
	"websocket/repo"
)

var providerSet = wire.NewSet(
	// connection
	connection.NewDB,
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

	// middleWare
	middleware.NewMiddleWare,

	// router
	router.NewRouter,

	wire.Struct(new(handler.Handler), "*"),
)

func Initialize() *router.Router {
	panic(wire.Build(providerSet))
}
