package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websocket/api/handler/ws"
	"websocket/config"
	"websocket/utils"
)

func ApiInfo(c *gin.Context) {
	info := map[string]interface{}{
		"app_name":      config.App.AppName,
		"online_user":   len(ws.GlobalHub.Users),
		"online_client": ws.GlobalHub.OnlineClient,
		"time":          utils.DateTime(),
	}

	c.JSON(http.StatusOK, info)
}
