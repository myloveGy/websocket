package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"websocket/api/ws"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket 处理websocket 信息
func WebSocket(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {

	var conn, err = upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	client := &ws.Client{Hub: hub, Conn: conn, Send: make(chan interface{}, 256)}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
