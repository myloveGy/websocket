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

	// 用户消息读取状态
	UserMessageUnread = 1 // 未读
	UserMessageRead   = 2 // 已经读取

	UserStatusActivate = 1 // 启用
	UserStatusDisabled = 2 // 体用

	// 消息类型
	MessageTypeTemp       = 1 // 临时消息
	MessageTypeNeedReplay = 2 // 必须回复消息
)
