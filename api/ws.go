package api

import (
	"log"
	"net/http"
	"time"

	"websocket/global"
	"websocket/utils"

	"github.com/gorilla/websocket"
)

// Message 发放的消息信息
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Time string `json:"time"`
}

var upgrade = websocket.Upgrader{
	HandshakeTimeout: 1 * time.Minute,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket 处理websocket 信息
func WebSocket(w http.ResponseWriter, r *http.Request) {
	var conn, err = upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer func() {
		log.Printf("send close \n")
		_ = conn.Close()
	}()

	_ = conn.SetReadDeadline(time.Now().Add(60 * time.Second))

	for {
		msg := &Message{}
		err := conn.ReadJSON(msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		log.Printf("recv: %v", msg)

		switch msg.Type {
		// 心跳检测
		case global.ClientHeartbeat:
			_ = conn.SetReadDeadline(time.Now().Add(35 * time.Second))
			_ = conn.WriteJSON(&Message{
				Type: global.ServerHeartbeat,
				Time: utils.DateTime(),
			})

			// 关闭链接
		case global.ClientClose: // 关闭链接
			log.Printf("close: %v", msg)
			_ = conn.Close()
			break
		case global.ClientAuth: // 连接上
			_ = conn.WriteJSON(&Message{
				Type: global.ServerAuth,
				Data: "已经建立链接，可以愉快的聊天了",
				Time: utils.DateTime(),
			})
		case global.ClientMessage:
			result, err := utils.GetHTTP(msg.Data)
			if err != nil || result.Code != 0 {
				log.Printf("机器人回复失败：%v\n", err)
				_ = conn.WriteJSON(&Message{
					Type: global.ServerMessage,
					Data: "机器人回复失败",
					Time: utils.DateTime(),
				})
				continue
			}

			log.Println(result)
			err = conn.WriteJSON(&Message{
				Type: global.ServerMessage,
				Data: result.Content,
				Time: utils.DateTime(),
			})

			if err != nil {
				log.Println("write:", err)
			}
		}
	}
}
