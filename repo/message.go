package repo

import (
	"github.com/jmoiron/sqlx"
	"time"
	"websocket/models"
)

type Message struct {
	db *sqlx.DB
}

func NewMessage(db *sqlx.DB) *Message {
	return &Message{db: db}
}

// Create 创建
func (m *Message) Create(model *models.Message) error {
	model.CreatedAt = time.Now()
	result, err := m.db.Exec(
		"INSERT INTO `message` (`app_id`, `type`, `content`, `created_at`) VALUES (?, ?, ?, ?)",
		model.AppId,
		model.Type,
		model.Content,
		model.CreatedAt,
	)
	if err != nil {
		return err
	}

	model.MessageId, err = result.LastInsertId()
	return err
}

// Delete 删除数据
func (m *Message) Delete(messageId int64) (int64, error) {
	result, err := m.db.Exec("DELETE FROM `message` WHERE `message_id` = ?", messageId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
