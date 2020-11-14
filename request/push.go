package request

// 推送用户消息
type UserParams struct {
	UserId  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	Type    int    `json:"type" binding:"required,oneof=1 2"`
}
