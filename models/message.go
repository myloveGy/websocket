package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Message struct {
	MessageId int64     `db:"message_id" json:"message_id"`
	AppId     int64     `db:"app_id" json:"app_id"`
	Type      int       `db:"type" json:"type"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at"`
}

// 创建
func (m *Message) Create(db *sqlx.DB) error {
	m.CreatedAt = time.Now()
	result, err := db.Exec(
		"INSERT INTO `message` (`app_id`, `type`, `content`, `created_at`) VALUES (?, ?, ?, ?)",
		m.AppId,
		m.Type,
		m.Content,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}

	m.MessageId, err = result.LastInsertId()
	return err
}

// 删除数据
func (m *Message) Delete(db *sqlx.DB) (int64, error) {
	result, err := db.Exec("DELETE FROM `message` WHERE `message_id` = ?", m.MessageId)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
