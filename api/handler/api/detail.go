package api

import (
	"github.com/gin-gonic/gin"

	"websocket/api/response"
	"websocket/cache"
	"websocket/config"
	"websocket/repo"
	"websocket/request"
	"websocket/service"
	"websocket/utils"
)

type Api struct {
	userCache *cache.UserCache
	userRepo  *repo.User
}

func NewApi(userCache *cache.UserCache, userRepo *repo.User) *Api {
	return &Api{
		userCache: userCache,
		userRepo:  userRepo,
	}
}

func (*Api) Detail(c *gin.Context) {
	info := make([]*request.AppItem, 0)
	for _, app := range service.GlobalHub.Apps {
		info = append(info, app.ToItem())
	}

	data := map[string]interface{}{
		"app_name":      config.App.AppName,
		"online_app":    len(service.GlobalHub.Apps),
		"online_client": service.GlobalHub.OnlineClient,
		"apps":          info,
		"time":          utils.DateTime(),
	}

	response.Success(c, data)
}
