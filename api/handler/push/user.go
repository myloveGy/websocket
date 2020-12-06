package push

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinxing-go/mysql"

	"websocket/api/response"
	"websocket/entity"
	"websocket/models"
	"websocket/request"
	"websocket/service"
	"websocket/service/api"
	"websocket/utils"
)

type Push struct {
	messageService *api.MessageService
}

func NewPush(messageService *api.MessageService) *Push {
	return &Push{messageService: messageService}
}

func (p *Push) User(context *gin.Context) {
	// 验证绑定数据
	params := &request.PushUserParams{}
	if isError, err := utils.BindAndValid(context, params); isError {
		response.InvalidParams(context, err.Error())
		return
	}

	// 获取应用消息
	app, err := p.getApp(context)
	if err != nil {
		response.BusinessError(context, err)
		return
	}

	// 添加消息
	messageList, err := p.messageService.CreateUserMessageList(app.Id, params.Message, []string{params.UserId})
	if err != nil {
		response.BusinessError(context, err)
		return
	}

	// 返回消息内容
	resp := &request.PushUserResponse{
		Online:  false,
		UserId:  params.UserId,
		Message: params.Message,
	}

	if value, ok := service.GlobalHub.Apps[app.Id]; ok {
		for _, message := range messageList {
			if user, ok1 := value.Users[params.UserId]; ok1 {
				for _, client := range user {
					client.Send <- &service.Message{
						Id:          message.Id,
						MessageId:   message.MessageId,
						MessageType: params.Type,
						Content:     params.Content,
						Type:        entity.SocketMessage,
						Time:        mysql.DateTime(),
					}
				}

				// 响应数据
				resp.Online = true
			}
		}
	}

	// 响应数据
	response.Success(context, resp)
}

func (p *Push) getApp(context *gin.Context) (*models.App, error) {
	// 获取应用消息
	app, ok := context.Get("app")
	if !ok {
		return nil, errors.New(entity.ErrAppNoTExists)
	}

	// 断言成App
	model, ok1 := app.(*models.App)
	if !ok1 {
		return nil, errors.New(entity.ErrApp)
	}

	return model, nil
}
