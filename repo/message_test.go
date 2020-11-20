package repo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"websocket/connection"
	"websocket/models"
)

func newTestMessageRepo() *Message {
	db := connection.NewDB()
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

	_, err1 := messageRepo.Delete(messageModel.MessageId)
	assert.NoError(t, err1)
}

func TestMessage_Delete(t *testing.T) {
	messageRepo := newTestMessageRepo()
	_, err := messageRepo.Delete(10000)
	assert.NoError(t, err)
}
