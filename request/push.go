package request

// 推送用户消息
type PushUserParams struct {
	UserId string `json:"user_id" binding:"required"`
	Message
}

// 推送用户消息返回
type PushUserResponse struct {
	Online bool   `json:"online"`
	UserId string `json:"user_id"`
	Message
}

// 推送多个用户消息
type PushManyParams struct {
	Users string `json:"users" binding:"required,min=2"`
	Message
}

type PushManyResponse struct {
	OnlineUsers  []string `json:"online_users"`
	OfflineUsers []string `json:"offline_users"`
	Message
}

type Message struct {
	Content string `json:"content" binding:"required"`
	Type    int    `json:"type" binding:"required,oneof=1 2"`
	GroupId string `json:"group_id,omitempty"`
}
