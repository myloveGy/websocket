package repo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"websocket/connection"
	"websocket/models"
)

func newTestMessageRepo() *Message {
	db := connection.NewMySQL()
	return NewMessage(db)
}

func TestMessage_Create(t *testing.T) {
	messageRepo := newTestMessageRepo()
	messageModel := &models.Message{
		AppId:     1,
		Type:      1,
		Content:   `{"username":"jinxing.liu","age":28}`,
		CreatedAt: time.Now(),
	}

	err := messageRepo.Create(messageModel)
	assert.NoError(t, err)

	_, err1 := messageRepo.Delete(&models.Message{MessageId: messageModel.MessageId})
	assert.NoError(t, err1)
}
