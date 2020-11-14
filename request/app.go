package request

type AppItem struct {
	AppId       string   `json:"app_id"`
	AppName     string   `json:"app_name"`
	OnlineUser  int      `json:"online_user"`
	OnlineGroup int      `json:"online_group"`
	OnlineUsers []string `json:"online_users"`
}
