package ws

import "sync"

// Hub 维护一组活动客户端，并发
type Hub struct {
	// 读写锁
	mu sync.RWMutex

	// 在线数链接数
	OnlineClient int64

	/**
	 * 连接用户信息
	 * 可能一个用户两个链接(手机端、网页端) {'user_id': [client1, client2]}
	 */
	Users map[string][]*Client

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
		Users:      make(map[string][]*Client),
	}
}

func (h *Hub) registerUser(c *Client) {

	h.mu.Lock()

	if h.Users == nil {
		h.Users = make(map[string][]*Client)
	}

	h.Users[c.UserId] = append(h.Users[c.UserId], c)

	h.mu.Unlock()
}

func (h *Hub) unRegisterUser(c *Client) {
	h.mu.Lock()
	var tmpUsers []*Client
	for _, v := range h.Users[c.UserId] {
		if v != c {
			tmpUsers = append(tmpUsers, v)
		}
	}

	h.Users[c.UserId] = tmpUsers
	h.mu.Unlock()
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			// 注册链接
			h.registerUser(client)
			h.OnlineClient += 1
		case client := <-h.Unregister:
			// 取消链接
			if _, ok := h.Users[client.UserId]; ok {
				h.unRegisterUser(client)
				close(client.Send)
				h.OnlineClient -= 1
			}
		}
	}
}
