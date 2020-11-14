package push

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/api/response"
	"websocket/global"
	"websocket/models"
	"websocket/request"
	"websocket/service"
	"websocket/utils"
)

type Push struct{}

func (p *Push) User(context *gin.Context) {
	// 验证绑定数据
	params := &request.UserParams{}
	if isError, err := utils.BindAndValid(context, params); isError {
		response.NewResponseError(context, "PushUserError", "请求参数错误:"+err.Error())
		return
	}

	// 获取应用消息
	app, ok := context.Get("app")
	appModel, ok1 := app.(*models.App)
	if !ok || !ok1 {
		response.NewResponseError(context, "PushUserError", "APP信息错误")
		return
	}

	if value, ok := service.GlobalHub.Apps[appModel.Id]; ok {
		if user, ok1 := value.Users[params.UserId]; ok1 {
			for _, client := range user {
				client.Send <- service.Message{
					Type:    global.SocketConnection,
					Content: params.Content,
					Time:    utils.DateTime(),
				}
			}

			// 响应数据
			context.JSON(http.StatusOK, map[string]interface{}{
				"online":     true,
				"user_id":    params.UserId,
				"content":    params.Content,
				"type":       params.Type,
				"created_at": utils.DateTime(),
			})
		}
	}

	// 添加数据入库
	m := &models.Message{
		Content: params.Content,
		Type:    params.Type,
		AppId:   appModel.Id,
	}

	// 添加消息
	if err := m.Create(global.DB); err != nil {
		response.NewResponseError(context, "PushUserError", "添加消息失败")
		return
	}

	// 添加用户消息
	mRead := &models.MessageRead{
		AppId:     m.AppId,
		MessageId: m.MessageId,
		UserId:    params.UserId,
	}

	if err := mRead.Create(global.DB); err != nil {
		response.NewResponseError(context, "PushUserError", "添加用户消息失败")
		return
	}

	// 响应数据
	response.Success(context, map[string]interface{}{
		"online":     false,
		"id":         mRead.Id,
		"message_id": mRead.MessageId,
		"user_id":    params.UserId,
		"content":    params.Content,
		"type":       params.Type,
		"created_at": utils.DateTime(),
	})
}
