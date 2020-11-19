package handler

import (
	"websocket/api/handler/api"
	"websocket/api/handler/push"
	"websocket/api/handler/user"
)

type Handler struct {
	Api  *api.Api
	Push *push.Push
	User *user.User
	Ws   *WS
}
