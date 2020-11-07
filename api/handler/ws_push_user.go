package handler

import (
	"net/http"
	"websocket/api/response"
	"websocket/utils"

	"github.com/gin-gonic/gin"
)

// 请求数据
type Params struct {
	SessionId string `json:"session_id" binding:"required"`
	Data      string `json:"data" binding:"required"`
	Type      int    `json:"type" binding:"required,oneof=1 2"`
}

func WsPushUser(context *gin.Context) {
	params := &Params{}
	if isError, err := utils.BindAndValid(context, params); isError {
		response.NewResponseError(context, "PushUserError", "请求参数错误:"+err.Error())
		return
	}

	context.JSON(http.StatusOK, params)
}
