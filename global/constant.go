package global

const (
	// 连接端发送消息类型
	ClientAuth           = "auth"            // 发送授权
	ClientMessage        = "message"         // 发送消息
	ClientHeartbeat      = "heartbeat"       // 心跳检测
	ClientClose          = "close"           // 主动关闭
	ClientMessageReceipt = "message receipt" // 消息回复

	// 服务端回复消息类型
	ServerMessage   = "message response"   // 消息回复
	ServerHeartbeat = "heartbeat response" // 心跳回复
	ServerAuth      = "auth response"      // 授权回复
)
