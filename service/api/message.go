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

// CreateUserMessageList 批量添加用户消息信息
func (m *MessageService) CreateUserMessageList(appId int64, params request.Message, users []string) ([]*models.MessageRead, error) {
	messageModels := make([]*models.MessageRead, 0)
	// 开启事务处理
	if err := m.messageRepo.Transaction(func(mysql *mysql.MySQl) error {
		// 第一步：添加消息信息
		model := &models.Message{
			AppId:   appId,
			Type:    params.Type,
			GroupId: params.GroupId,
			Content: params.Content,
		}

		if err := m.messageRepo.Create(model); err != nil {
			return errors.New(entity.ErrCreateMessage)
		}

		for _, userId := range users {
			message := &models.MessageRead{
				MessageId: model.MessageId,
				AppId:     model.AppId,
				UserId:    userId,
				Status:    entity.UserMessageUnread,
			}

			if err := mysql.Create(message); err != nil {
				return errors.New(entity.ErrCreateUserMessage)
			}

			messageModels = append(messageModels, message)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return messageModels, nil
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
