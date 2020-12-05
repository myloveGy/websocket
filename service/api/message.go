package api

import (
	"errors"
	"github.com/jinxing-go/mysql"
	"websocket/entity"

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

func (m *MessageService) Create(appId int64, params request.Message) (*models.Message, error) {
	model := &models.Message{
		AppId:   appId,
		Type:    params.Type,
		GroupId: params.GroupId,
		Content: params.Content,
	}

	if err := m.messageRepo.Create(model); err != nil {
		return nil, errors.New(entity.ErrCreateMessage)
	}

	return model, nil
}

func (m *MessageService) CreateUserMessage(userIds []string, message *models.Message) error {
	// 开启事务处理
	if err := m.messageRepo.Transaction(func(mysql *mysql.MySQl) error {
		// 添加用户消息
		for _, userId := range userIds {
			if err := mysql.Create(&models.MessageRead{
				MessageId: message.MessageId,
				AppId:     message.AppId,
				UserId:    userId,
				Status:    entity.UserMessageUnread,
			}); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return errors.New(entity.ErrCreateUserMessage)
	}

	return nil
}
