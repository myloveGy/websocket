package admin

type AppSearch struct {
	AppId   string `json:"app_id"`
	AppName string `json:"app_name"`
	Status  int    `json:"status"`
	AdminSearch
}

type AppCreate struct {
	AppSecret string `json:"app_secret" binding:"required,min=16"`
	AppName   string `json:"app_name" binding:"required,min=2"`
}

type AppIdStruct struct {
	Id int64 `json:"id" binding:"required,min=1"`
}

type AppUpdate struct {
	AppIdStruct
	AppCreate
}
