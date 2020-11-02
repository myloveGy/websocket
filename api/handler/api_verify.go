package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websocket/api/handler/ws"
	"websocket/config"
	"websocket/utils"
)

func ApiVerify(c *gin.Context) {
	info := map[string]interface{}{
		"app_name": config.App.AppName,
		"clients":  len(ws.GlobalHub.Clients),
		"time":     utils.DateTime(),
	}

	c.JSON(http.StatusOK, info)
}
