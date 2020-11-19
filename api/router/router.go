package router

import (
	"websocket/api/handler"
	"websocket/api/middleware"
)

type Router struct {
	middleware *middleware.MiddleWare
	handler    *handler.Handler
}

func NewRouter(m *middleware.MiddleWare, h *handler.Handler) *Router {
	return &Router{
		middleware: m,
		handler:    h,
	}
}
