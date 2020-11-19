package middleware

import (
	"websocket/cache"
	"websocket/repo"
)

type MiddleWare struct {
	userCache *cache.UserCache
	appRepo   *repo.App
}

func NewMiddleWare(userCache *cache.UserCache, appRepo *repo.App) *MiddleWare {
	return &MiddleWare{
		userCache: userCache,
		appRepo:   appRepo,
	}
}
