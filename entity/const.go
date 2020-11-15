package entity

const (
	// Socket消息类型
	SocketConnection     = "connection"    // 已经连接
	SocketMessage        = "message"       // 发送消息
	SocketHeartbeat      = "heartbeat"     // 心跳检测
	SocketClose          = "close"         // 主动关闭
	SocketMessageReceipt = "reply_message" // 消息回复

	// 应用状态
	AppStatusActivate = 1 // 状态启用
	AppStatusDisabled = 2 // 停用
)
