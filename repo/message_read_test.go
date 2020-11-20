package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"websocket/connection"
	"websocket/models"
)

func newTestMessageRead() *MessageRead {
	db := connection.NewDB()
	return NewMessageRead(db)
}

func TestMessageRead_Create(t *testing.T) {
	messageReadRepo := newTestMessageRead()
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

		intRow, err1 := messageReadRepo.Delete(messageReadModel.Id)
		assert.NoError(t, err1)
		assert.Equal(t, int64(1), intRow)
	})
}

func TestMessageRead_FindAll(t *testing.T) {
	messageReadRepo := newTestMessageRead()
	_, err := messageReadRepo.FindAll(1, "1", 1)
	assert.NoError(t, err)
}

func TestMessageRead_UpdateStatus(t *testing.T) {
	messageReadRepo := newTestMessageRead()
	_, err := messageReadRepo.UpdateStatus(1, 2)
	assert.NoError(t, err)
}
