package admin

type UserSearch struct {
	UserId    string `json:"user_id"`
	AppId     int64  `json:"app_id"`
	Username  string `json:"username"`
	Status    int    `json:"status"`
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

type UserId struct {
	UserId int64 `json:"user_id" binding:"required,min=1"`
}
