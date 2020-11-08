package handler

import (
	"net/http"
	"websocket/api/handler/ws"
	"websocket/api/response"
	"websocket/global"
	"websocket/models"
	"websocket/utils"

	"github.com/gin-gonic/gin"
)

// 请求数据
type Params struct {
	UserId  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
	Type    int    `json:"type" binding:"required,oneof=1 2"`
}

func WsPushUser(context *gin.Context) {
	// 验证绑定数据
	params := &Params{}
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

	if value, ok := ws.GlobalHub.Users[params.UserId]; ok {
		for _, client := range value {
			client.Send <- ws.Message{
				Type: global.ClientMessage,
				Data: params.Content,
				Time: utils.DateTime(),
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
	} else {
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
		context.JSON(http.StatusOK, map[string]interface{}{
			"online":     false,
			"id":         mRead.Id,
			"message_id": mRead.MessageId,
			"user_id":    params.UserId,
			"content":    params.Content,
			"type":       params.Type,
			"created_at": utils.DateTime(),
		})
	}
}
