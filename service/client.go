package service

import (
	"encoding/json"
	"fmt"
	"github.com/jinxing-go/mysql"
	"log"
	"time"
	"websocket/repo"

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
	Send chan *Message

	MessageReadRepo *repo.MessageRead
}

// Message 发放的消息信息
type Message struct {
	Id          int64  `json:"id,omitempty"`
	Type        string `json:"type"`
	MessageId   int64  `json:"-"`
	MessageType int    `json:"-"`
	Content     string `json:"content,omitempty"`
	Time        string `json:"time"`
}

type ReplyMessage struct {
	Id int64 `json:"id"`
}

type ReplyMessageResponse struct {
	Id      int64  `json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
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

		var responseMessage = &Message{
			Type: msg.Type,
		}

		switch msg.Type {
		// 关闭链接
		case entity.SocketClose:
			log.Printf("close: %v", msg)
			_ = c.Conn.Close()
			break
		case entity.SocketMessageReceipt:
			m := &ReplyMessage{}
			// 解析消息信息
			if err := json.Unmarshal([]byte(msg.Content), m); err != nil {
				str, _ := json.Marshal(&ReplyMessageResponse{
					Status:  "error",
					Message: fmt.Sprintf("%s(%s)", entity.ErrDecodeReplyMessage, msg.Content),
				})

				responseMessage.Content = string(str)
				break
			}

			// 修改消息信息
			if _, err := c.MessageReadRepo.UpdateStatus(m.Id, entity.UserMessageRead); err != nil {
				str, _ := json.Marshal(&ReplyMessageResponse{
					Id:      m.Id,
					Status:  "error",
					Message: entity.ErrUpdateUserMessage,
				})

				responseMessage.Content = string(str)
				break
			}

			str, _ := json.Marshal(&ReplyMessageResponse{
				Id:      m.Id,
				Status:  "success",
				Message: "OK",
			})

			responseMessage.Content = string(str)

		case entity.SocketMessage:
			m := &struct {
				Source  string `json:"source"`
				Content string `json:"content"`
			}{}

			err := json.Unmarshal([]byte(msg.Content), m)
			if err != nil {
				responseMessage.Content = "你发的什么: " + err.Error()
				break
			}

			result, err := utils.GetHTTP(m.Content)
			if err != nil || result.Code != 0 {
				log.Printf("机器人回复失败：%v\n", err)
				responseMessage.Content = "机器人回复失败"
				break
			}

			responseMessage.Content = result.Content
		}

		if responseMessage.Type != "" {
			responseMessage.Time = mysql.DateTime()
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

			// 如果消息为临时消息，需要标记为已经读取
			if message.MessageType == entity.MessageTypeTemp && message.Id > 0 && c.MessageReadRepo != nil {
				c.MessageReadRepo.UpdateStatus(message.Id, entity.UserMessageRead)
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
