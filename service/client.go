package service

import (
	"encoding/json"
	"github.com/jinxing-go/mysql"
	"log"
	"time"

	"websocket/models"

	"github.com/gorilla/websocket"

	"websocket/entity"
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
	App *models.App

	// 用户信息
	UserId string

	// 集合
	Hub *Hub

	// 连接
	Conn *websocket.Conn

	// 发送消息
	Send chan interface{}
}

// Message 发放的消息信息
type Message struct {
	Type    string `json:"type"`
	Content string `json:"content,omitempty"`
	Time    string `json:"time"`
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
		case entity.SocketHeartbeat:
			responseMessage = Message{
				Type: entity.SocketHeartbeat,
				Time: mysql.DateTime(),
			}

			// 关闭链接
		case entity.SocketClose:
			log.Printf("close: %v", msg)
			_ = c.Conn.Close()
			break
		case entity.SocketMessage:
			m := &struct {
				Source  string `json:"source"`
				Content string `json:"content"`
			}{}

			err := json.Unmarshal([]byte(msg.Content), m)
			if err != nil {
				responseMessage = Message{
					Type:    entity.SocketMessage,
					Content: "你发的什么: " + err.Error(),
					Time:    mysql.DateTime(),
				}

				break
			}

			result, err := utils.GetHTTP(m.Content)
			if err != nil || result.Code != 0 {
				log.Printf("机器人回复失败：%v\n", err)

				responseMessage = Message{
					Type:    entity.SocketMessage,
					Content: "机器人回复失败",
					Time:    mysql.DateTime(),
				}

				break
			}

			responseMessage = Message{
				Type:    entity.SocketMessage,
				Content: result.Content,
				Time:    mysql.DateTime(),
			}
		}

		if responseMessage.Type != "" {
			c.Send <- responseMessage
		}
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
