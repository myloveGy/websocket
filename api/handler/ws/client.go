package ws

import (
	"log"
	"time"

	"github.com/gorilla/websocket"

	"websocket/global"
	"websocket/utils"
)

const (
	// 等待客户端写入消息时间
	writeWait = 10 * time.Second

	// 等待读取下一条消息时间
	pongWait = 60 * time.Second

	// 发送ping到peer与此周期。必须小于pongWait。
	pingPeriod = (pongWait * 9) / 10

	// peer允许的最大消息大小。
	maxMessageSize = 512
)

type Client struct {
	// 集合
	Hub *Hub

	// 连接
	Conn *websocket.Conn

	// 发送消息
	Send chan interface{}
}

// Message 发放的消息信息
type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Time string `json:"time"`
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		_ = c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		msg := &Message{}
		err := c.Conn.ReadJSON(msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var responseMessage Message

		switch msg.Type {
		// 心跳检测
		case global.ClientHeartbeat:
			responseMessage = Message{
				Type: global.ServerHeartbeat,
				Time: utils.DateTime(),
			}

			// 关闭链接
		case global.ClientClose: // 关闭链接
			log.Printf("close: %v", msg)
			_ = c.Conn.Close()
			break
		case global.ClientAuth: // 连接上
			responseMessage = Message{
				Type: global.ServerAuth,
				Data: "认证成功，已经建立链接",
				Time: utils.DateTime(),
			}
		case global.ClientMessage:
			result, err := utils.GetHTTP(msg.Data)
			if err != nil || result.Code != 0 {
				log.Printf("机器人回复失败：%v\n", err)

				responseMessage = Message{
					Type: global.ServerMessage,
					Data: "机器人回复失败",
					Time: utils.DateTime(),
				}

				break
			}

			responseMessage = Message{
				Type: global.ServerMessage,
				Data: result.Content,
				Time: utils.DateTime(),
			}
		}

		c.Send <- responseMessage
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.Conn.WriteJSON(message)
			if err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}