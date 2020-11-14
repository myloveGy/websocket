package service

import (
	"sync"
	"websocket/models"
)

// Hub 维护一组活动客户端，并发
type Hub struct {
	// 读写锁
	mu sync.RWMutex

	// 在线数链接数
	OnlineClient int64

	Apps map[int64]*App

	// 注册来自客户端的请求。
	Register chan *Client

	// 取消客户端的注册请求
	Unregister chan *Client
}

var GlobalHub *Hub

func init() {
	GlobalHub = NewHub()
}

// NewHub 创建Hub
func NewHub() *Hub {
	return &Hub{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Apps:       make(map[int64]*App),
	}
}

func (h *Hub) CreateApp(app *models.App) {

	h.mu.Lock()

	if h.Apps == nil {
		h.Apps = make(map[int64]*App)
	}

	if _, ok := h.Apps[app.Id]; !ok {
		h.Apps[app.Id] = NewApp(app)
	}

	h.mu.Unlock()
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// 注册链接
			h.CreateApp(client.App)
			if v, ok := h.Apps[client.App.Id]; ok {
				v.register(client)
			}

			h.OnlineClient += 1
		case client := <-h.Unregister:
			if v, ok := h.Apps[client.App.Id]; ok {
				v.unRegister(client)
				close(client.Send)
				h.OnlineClient -= 1
			}
		}
	}
}
