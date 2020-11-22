package push

import (
	"github.com/gin-gonic/gin"
	"github.com/jinxing-go/mysql"
	"websocket/api/response"
	"websocket/entity"
	"websocket/models"
	"websocket/repo"
	"websocket/request"
	"websocket/service"
	"websocket/utils"
)

type Push struct {
	messageRepo     *repo.Message
	messageReadRepo *repo.MessageRead
}

func NewPush(messageRepo *repo.Message, messageReadRepo *repo.MessageRead) *Push {
	return &Push{messageRepo: messageRepo, messageReadRepo: messageReadRepo}
}

func (p *Push) User(context *gin.Context) {
	// 验证绑定数据
	params := &request.UserParams{}
	if isError, err := utils.BindAndValid(context, params); isError {
		response.InvalidParams(context, err.Error())
		return
	}

	// 获取应用消息
	app, ok := context.Get("app")
	appModel, ok1 := app.(*models.App)
	if !ok || !ok1 {
		response.BusinessError(context, "APP信息错误")
		return
	}

	if value, ok := service.GlobalHub.Apps[appModel.Id]; ok {
		if user, ok1 := value.Users[params.UserId]; ok1 {
			for _, client := range user {
				client.Send <- service.Message{
					Type:    entity.SocketConnection,
					Content: params.Content,
					Time:    mysql.DateTime(),
				}
			}

			// 响应数据
			response.Success(context, map[string]interface{}{
				"online":     true,
				"user_id":    params.UserId,
				"content":    params.Content,
				"type":       params.Type,
				"created_at": mysql.DateTime(),
			})

			return
		}
	}

	// 添加数据入库
	m := &models.Message{
		Content: params.Content,
		Type:    params.Type,
		AppId:   appModel.Id,
	}

	// 添加消息
	if err := p.messageRepo.Create(m); err != nil {
		response.SystemError(context, "添加消息失败")
		return
	}

	// 添加用户消息
	mRead := &models.MessageRead{
		AppId:     m.AppId,
		MessageId: m.MessageId,
		UserId:    params.UserId,
	}

	if err := p.messageReadRepo.Create(mRead); err != nil {
		response.SystemError(context, "添加用户消息失败")
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
		"created_at": mysql.DateTime(),
	})
}
