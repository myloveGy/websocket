package middleware

import (
	"github.com/gin-gonic/gin"

	"websocket/api/response"
)

func (m *MiddleWare) AccessToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 优先从header 获取，然后从请求参数中获取
		accessToken := context.Request.Header.Get("Access-Token")
		if accessToken == "" {
			accessToken = context.Query("access_token")
		}

		// accessToken 信息为空
		if accessToken == "" {
			response.Unauthorized(context)
			return
		}

		// redis 获取用户信息
		user, err := m.userCache.Get(accessToken)
		if err != nil {
			response.NotLogin(context, err)
			return
		}

		// 查询应用信息
		app, err := m.appRepo.FindById(user.AppId)
		if err != nil {
			response.BusinessError(context, "应用信息不存在")
			return
		}

		context.Set("user", user)
		context.Set("app", app)
		context.Next()
	}
}
