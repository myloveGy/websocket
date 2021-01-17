package app

import (
	"github.com/gin-gonic/gin"

	"websocket/api/response"
	"websocket/request/admin"
	"websocket/utils"
)

func (a *App) Delete(c *gin.Context) {
	// 解析参数
	params := &admin.AppIdStruct{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	// 删除数据
	if err := a.appService.Delete(params); err != nil {
		response.BusinessError(c, err)
		return
	}

	response.Success(c, params)
	return
}
