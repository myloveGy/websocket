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

func NewHandler(api *api.Api, push *push.Push, user *user.User, ws *WS) *Handler {
	return &Handler{
		Api:  api,
		Push: push,
		User: user,
		Ws:   ws,
	}
}
