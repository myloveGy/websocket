package admin

type AdminSearch struct {
	SortField string `json:"sort_field"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
}

type UserSearch struct {
	UserId   string `json:"user_id"`
	AppId    int64  `json:"app_id"`
	Username string `json:"username"`
	Status   int    `json:"status"`
	AdminSearch
}

type UserId struct {
	UserId int64 `json:"user_id" binding:"required,min=1"`
}

type UserCreate struct {
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password"`
}

type UserUpdate struct {
	UserId
	UserCreate
}
