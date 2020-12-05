package push

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinxing-go/mysql"

	"websocket/api/response"
	"websocket/entity"
	"websocket/request"
	"websocket/service"
	"websocket/utils"
)

func (p *Push) Many(context *gin.Context) {
	// 验证绑定数据
	params := &request.PushManyParams{}
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

	users := make([]string, 0)
	if err := json.Unmarshal([]byte(params.Users), &users); err != nil {
		response.BusinessError(context, "Users信息错误")
		return
	}

	// 添加消息
	message, err := p.messageService.Create(app.Id, params.Message)
	if err != nil {
		response.BusinessError(context, err)
		return
	}

	// 返回消息内容
	resp := &request.PushManyResponse{
		Message:      params.Message,
		OnlineUsers:  make([]string, 0),
		OfflineUsers: make([]string, 0),
	}

	if value, ok := service.GlobalHub.Apps[app.Id]; ok {
		for _, userId := range users {
			if user, ok1 := value.Users[userId]; ok1 {
				for _, client := range user {
					client.Send <- &service.Message{
						MessageId:   message.MessageId,
						Type:        entity.SocketMessage,
						MessageType: params.Type,
						Content:     params.Content,
						Time:        mysql.DateTime(),
					}
				}

				resp.OnlineUsers = append(resp.OnlineUsers, userId)
			} else {
				resp.OfflineUsers = append(resp.OfflineUsers, userId)
			}
		}
	} else {
		resp.OfflineUsers = users
	}

	if params.Type == entity.MessageTypeTemp {
		users = resp.OfflineUsers
	}

	// 添加消息
	if err := p.messageService.CreateUserMessage(users, message); err != nil {
		response.SystemError(context, err)
		return
	}

	// 响应数据
	response.Success(context, resp)
}
