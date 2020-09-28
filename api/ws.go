package api

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"websocket/global"
	"websocket/utils"
)

// Message 发放的消息信息
type Message struct {
	Type   string `json:"type"`
	Data   string `json:"data"`
	Source string `json:"source"`
}

var upgrade = websocket.Upgrader{
	HandshakeTimeout: 1 * time.Minute,
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
		case global.WsHeartbeat:
			_ = conn.SetReadDeadline(time.Now().Add(35 * time.Second))
			// 关闭链接
		case global.WsClose: // 关闭链接
			log.Printf("close: %v", msg)
			_ = conn.Close()
			break
		case global.WsOpen: // 连接上
			_ = conn.WriteJSON(&Message{
				Type:   "system",
				Data:   "已经建立链接，可以愉快的聊天了",
				Source: "system",
			})
		case global.WsMessage:
			result, err := utils.GetHTTP(msg.Data)
			if err != nil || result.Code != 0 {
				log.Printf("机器人回复失败：%v\n", err)
				_ = conn.WriteJSON(&Message{
					Type:   "system",
					Data:   "机器人回复失败",
					Source: "system",
				})
				continue
			}

			log.Println(result)
			err = conn.WriteJSON(&Message{
				Type:   "system",
				Data:   result.Content,
				Source: "system",
			})

			if err != nil {
				log.Println("write:", err)
			}
		}
	}
}
