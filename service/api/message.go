package api

import (
	"github.com/jinxing-go/mysql"

	"websocket/models"
	"websocket/repo"
	"websocket/request"
)

type MessageService struct {
	messageRepo     *repo.Message
	messageReadRepo *repo.MessageRead
}

func NewMessageService(messageRepo *repo.Message, messageReadRepo *repo.MessageRead) *MessageService {
	return &MessageService{
		messageRepo:     messageRepo,
		messageReadRepo: messageReadRepo,
	}
}

func (m *MessageService) BatchCreateUserMessage(appId int64, userIds []string, message request.Message) error {
	// 开启事务处理
	return m.messageRepo.Transaction(func(mysql *mysql.MySQl) error {
		// 新增消息内容
		mModel := &models.Message{
			Content: message.Content,
			Type:    message.Type,
			AppId:   appId,
		}

		if err := mysql.Create(mModel); err != nil {
			return err
		}

		// 添加用户消息
		for _, userId := range userIds {
			if err := mysql.Create(&models.MessageRead{
				MessageId: mModel.MessageId,
				AppId:     mModel.AppId,
				UserId:    userId,
				Status:    1,
			}); err != nil {
				return err
			}
		}

		return nil
	})
}
