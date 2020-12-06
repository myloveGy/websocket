package handler

import (
	adminUser "websocket/api/handler/admin/user"
	"websocket/api/handler/api"
	"websocket/api/handler/push"
	"websocket/api/handler/user"
)

type Handler struct {
	Api       *api.Api
	Push      *push.Push
	User      *user.User
	Ws        *WS
	AdminUser *adminUser.User
}
