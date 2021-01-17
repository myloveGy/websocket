package handler

import (
	adminApp "websocket/api/handler/admin/app"
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
	AdminApp  *adminApp.App
}
