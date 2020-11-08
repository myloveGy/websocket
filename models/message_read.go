package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type MessageRead struct {
	Id        int64     `db:"id" json:"id"`
	MessageId int64     `db:"message_id" json:"message_id"`
	AppId     int       `db:"app_id" json:"app_id"`
	UserId    string    `db:"user_id" json:"user_id"`
	GroupId   string    `db:"group_id" json:"group_id"`
	Status    int       `db:"status" json:"status"`
	Content   string    `db:"content" json:"content"`
	Type      int       `db:"type" json:"type"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (m *MessageRead) FindAll(db *sqlx.DB) ([]*MessageRead, error) {
	list := make([]*MessageRead, 0)
	err := db.Select(&list, "SELECT "+
		"`message_read`.*, `message`.`content`, `message`.`type` "+
		"FROM `message_read` "+
		"INNER JOIN `message` ON (`message_read`.`message_id` = `message`.`message_id`) "+
		"WHERE  `message_read`.`app_id` = ? AND `message_read`.`user_id` = ? AND `message_read`.`status` = ? "+
		"ORDER BY `created_at` ASC", m.AppId, m.UserId, m.Status)
	return list, err
}

// 创建
func (m *MessageRead) Create(db *sqlx.DB) error {
	result, err := db.Exec(
		"INSERT INTO `message_read` (`message_id`, `app_id`, `user_id`, `group_id`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?)",
		m.MessageId,
		m.AppId,
		m.UserId,
		m.GroupId,
		m.CreatedAt,
		m.UpdatedAt,
	)
	if err != nil {
		return err
	}

	m.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

// 删除数据
func (m *MessageRead) Delete(db *sqlx.DB) (int64, error) {
	result, err := db.Exec("DELETE FROM `message_read` WHERE `id` = ?", m.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (m *MessageRead) UpdateStatus(db *sqlx.DB) (int64, error) {
	result, err := db.Exec("UPDATE `message_read` SET `status` = ? WHERE `id` = ?", m.Status, m.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
