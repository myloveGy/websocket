package repo

import (
	"testing"

	"github.com/jinxing-go/mysql"
	"github.com/stretchr/testify/assert"
	"websocket/models"
)

func newTestMessageRead(t *testing.T) *MessageRead {
	mySQL := mysql.NewTestMySQL(t, "../testdata/websocket.sql")
	return NewMessageRead(mySQL)
}

func TestMessageRead_Create(t *testing.T) {
	messageReadRepo := newTestMessageRead(t)
	// 创建成功
	t.Run("创建成功", func(t *testing.T) {
		messageReadModel := &models.MessageRead{
			AppId:     1,
			MessageId: 1,
			UserId:    "1",
			GroupId:   "",
		}

		err := messageReadRepo.Create(messageReadModel)
		assert.NoError(t, err)

		intRow, err1 := messageReadRepo.Delete(&models.MessageRead{Id: messageReadModel.Id})
		assert.NoError(t, err1)
		assert.Equal(t, int64(1), intRow)
	})
}

func TestMessageRead_FindAll(t *testing.T) {
	messageReadRepo := newTestMessageRead(t)
	_, err := messageReadRepo.FindAll(1, "1", 1)
	assert.NoError(t, err)
}

func TestMessageRead_UpdateStatus(t *testing.T) {
	messageReadRepo := newTestMessageRead(t)
	_, err := messageReadRepo.UpdateStatus(1, 2)
	assert.NoError(t, err)
}
