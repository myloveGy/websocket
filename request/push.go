package request

// 推送用户消息
type UserParams struct {
	UserId string `json:"user_id" binding:"required"`
	Message
}

type Message struct {
	Content string `json:"content" binding:"required"`
	Type    int    `json:"type" binding:"required,oneof=1 2"`
	GroupId string `json:"group_id"`
}
