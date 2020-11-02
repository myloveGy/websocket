package ws

// Hub 维护一组活动客户端，并发
type Hub struct {
	// 注册的连接
	Clients map[*Client]bool

	// 来自客户端的入站消息
	Broadcast chan interface{}

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
		Broadcast:  make(chan interface{}),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
