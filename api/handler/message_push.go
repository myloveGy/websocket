package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	SessionId string `form:"session_id" json:"session_id" binding:"required"`
	UserName  string `form:"username" json:"username" binding:"required"`
}

func MessagePush(context *gin.Context) {
	params := &Params{}
	if err := context.ShouldBind(params); err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    "Middleware Request error",
			"message": "请求参数错误",
		})

		context.Abort()
		return
	}

	context.JSON(http.StatusOK, params)
}
