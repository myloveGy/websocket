package app

import (
	"github.com/gin-gonic/gin"
	"websocket/api/response"
	"websocket/config"
	"websocket/request/admin"
	"websocket/service/api"
	"websocket/utils"
)

type App struct {
	appService *api.AppService
}

func NewApp(appService *api.AppService) *App {
	return &App{appService: appService}
}

func (a *App) List(c *gin.Context) {
	// 解析参数
	params := &admin.AppSearch{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	if params.Page <= 0 {
		params.Page = 1
	}

	if params.PageSize == 0 {
		params.PageSize = config.App.DefaultPageSize
	}

	// 查询数据
	users, total, err := a.appService.List(params)
	if err != nil {
		response.BusinessError(c, err)
		return
	}

	response.PageSuccess(c, users, total, params.Page)
	return
}
