package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websocket/api/handler/ws"
	"websocket/config"
	"websocket/utils"
)

type appInfo struct {
	AppId       string `json:"app_id"`
	AppName     string `json:"app_name"`
	OnlineUser  int    `json:"online_user"`
	OnlineGroup int    `json:"online_group"`
}

func ApiInfo(c *gin.Context) {
	info := make([]appInfo, 0)
	for _, app := range ws.GlobalHub.Apps {
		info = append(info, appInfo{
			AppId:       app.AppId,
			AppName:     app.AppName,
			OnlineUser:  len(app.Users),
			OnlineGroup: len(app.Groups),
		})
	}

	data := map[string]interface{}{
		"app_name":      config.App.AppName,
		"online_app":    len(ws.GlobalHub.Apps),
		"online_client": ws.GlobalHub.OnlineClient,
		"apps":          info,
		"time":          utils.DateTime(),
	}

	c.JSON(http.StatusOK, data)
}
