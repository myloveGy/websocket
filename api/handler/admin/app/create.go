package app

import (
	"github.com/gin-gonic/gin"
	"websocket/api/response"
	"websocket/request/admin"
	"websocket/utils"
)

func (a *App) Create(c *gin.Context) {
	// 解析参数
	params := &admin.AppCreate{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	// 创建数据
	if app, err := a.appService.Create(params); err != nil {
		response.BusinessError(c, err)
	} else {
		response.Success(c, app)
	}

	return
}
