package handler

import (
	"net/http"
	"websocket/api/response"

	"github.com/gin-gonic/gin"
)

// 请求数据
type Params struct {
	SessionId string `form:"session_id" json:"session_id" binding:"required"`
	Data      string `form:"data" json:"data" binding:"required"`
}

func WsPushUser(context *gin.Context) {
	params := &Params{}
	if err := context.ShouldBind(params); err != nil {
		response.NewResponseError(context, "PushUserError", "请求参数错误")
		return
	}

	context.JSON(http.StatusOK, params)
}
