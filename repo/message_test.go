package repo

import (
	"testing"
	"time"

	"github.com/jinxing-go/mysql"
	"github.com/stretchr/testify/assert"

	"websocket/models"
)

func newTestMessageRepo(t *testing.T) *Message {
	mySQL := mysql.NewTestMySQL(t, "../testdata/websocket.sql")
	return NewMessage(mySQL)
}

func TestMessage_Create(t *testing.T) {
	messageRepo := newTestMessageRepo(t)
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
