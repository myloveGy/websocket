package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"websocket/api/handler/ws"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket 处理websocket 信息
func WebSocket(c *gin.Context) {

	var conn, err = upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := &ws.Client{Hub: ws.GlobalHub, Conn: conn, Send: make(chan interface{}, 256)}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
