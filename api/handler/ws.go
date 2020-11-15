package handler

import (
	"net/http"
	"websocket/api/response"
	"websocket/entity"
	"websocket/repo"
	"websocket/service"
	"websocket/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WS struct {
	appRepo *repo.App
}

func NewWs(app *repo.App) *WS {
	return &WS{appRepo: app}
}

// WebSocket 处理websocket 信息
func (w *WS) WebSocket(c *gin.Context) {
	// 验证参数
	appId := c.Param("app_id")
	if appId == "" {
		response.InvalidParams(c, "传递参数错误")
		return
	}

	// 查询应用信息
	app, err := w.appRepo.FindByAppId(c.Param("app_id"))
	if err != nil {
		response.InvalidParams(c, "应用信息不存在: "+err.Error())
		return
	}

	// 验证应用状态
	if app.Status != entity.AppStatusActivate {
		response.BusinessError(c, "应用信息已经被停用")
		return
	}

	user_id := c.Query("user_id")
	if user_id == "" {
		response.InvalidParams(c, "user_id不能为空")
		return
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.SystemError(c, "upgrade error")
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
		Type:    entity.SocketConnection,
		Content: "已经建立链接",
		Time:    utils.DateTime(),
	}

	go client.WritePump()
	go client.ReadPump()
}
