package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinxing-go/mysql"
	"net/http"
	"websocket/api/response"
	"websocket/entity"
	"websocket/repo"
	"websocket/service"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WS struct {
	appRepo         *repo.App
	messageReadRepo *repo.MessageRead
}

func NewWs(app *repo.App, messageReadRepo *repo.MessageRead) *WS {
	return &WS{appRepo: app, messageReadRepo: messageReadRepo}
}

// WebSocket 处理websocket 信息
func (w *WS) WebSocket(c *gin.Context) {
	// 验证参数
	appId := c.Param("app_id")
	if appId == "" {
		response.InvalidParams(c)
		return
	}

	// 查询应用信息
	app, err := w.appRepo.FindByAppId(c.Param("app_id"))
	if err != nil {
		response.InvalidParams(c, entity.ErrAppNoTExists)
		return
	}

	// 验证应用状态
	if app.Status != entity.AppStatusActivate {
		response.BusinessError(c, entity.ErrAppDisable)
		return
	}

	userId := c.Query("user_id")
	if userId == "" {
		response.InvalidParams(c, entity.ErrUserNotEmpty)
		return
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.SystemError(c, "upgrade error")
		return
	}

	client := &service.Client{
		Hub:             service.GlobalHub,
		Conn:            conn,
		Send:            make(chan *service.Message, 256),
		UserId:          userId,
		App:             app,
		MessageReadRepo: w.messageReadRepo,
	}

	client.Hub.Register <- client
	client.Send <- &service.Message{
		Type:    entity.SocketConnection,
		Content: "已经建立链接",
		Time:    mysql.DateTime(),
	}

	// 查询用户信息
	if messageList, err := w.messageReadRepo.FindAll(app.Id, userId, entity.UserMessageUnread); err == nil {
		for _, v := range messageList {
			client.Send <- &service.Message{
				Id:          v.Id,
				MessageId:   v.MessageId,
				MessageType: v.Type,
				Type:        entity.SocketMessage,
				Content:     v.Content,
				Time:        mysql.DateTime(),
			}
		}
	}

	go client.WritePump()
	go client.ReadPump()
}
