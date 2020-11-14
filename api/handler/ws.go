package handler

import (
	"net/http"
	"websocket/api/response"
	"websocket/global"
	"websocket/service"
	"websocket/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"websocket/models"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket 处理websocket 信息
func WebSocket(c *gin.Context) {
	// 验证参数
	appId := c.Param("app_id")
	if appId == "" {
		response.NewResponseError(c, "WebSocketError", "传递参数错误")
		return
	}

	// 查询应用信息
	app := &models.App{AppId: c.Param("app_id")}
	if err := app.Find(); err != nil {
		response.NewResponseError(c, "WebSocketError", "应用信息不存在: "+err.Error())
		return
	}

	// 验证应用状态
	if app.Status != global.AppStatusActivate {
		response.NewResponseError(c, "WebSocketError", "应用信息已经被停用")
		return
	}

	user_id := c.Query("user_id")
	if user_id == "" {
		response.NewResponseError(c, "WebSocketError", "user_id不能为空")
		return
	}

	var conn, err = upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.NewResponseError(c, "WebSocketError", "upgrade error")
		return
	}

	client := &service.Client{
		Hub:    service.GlobalHub,
		Conn:   conn,
		Send:   make(chan interface{}, 256),
		UserId: user_id,
		App:    app,
	}

	client.Hub.Register <- client
	client.Send <- service.Message{
		Type:    global.SocketConnection,
		Content: "已经建立链接",
		Time:    utils.DateTime(),
	}

	go client.WritePump()
	go client.ReadPump()
}
