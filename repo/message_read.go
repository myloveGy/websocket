package repo

import (
	"github.com/jmoiron/sqlx"
	"time"
	"websocket/models"
)

type MessageRead struct {
	db *sqlx.DB
}

func NewMessageRead(db *sqlx.DB) *MessageRead {
	return &MessageRead{db: db}
}

func (m *MessageRead) FindAll(appId int64, userId string, status int) ([]*models.MessageRead, error) {
	list := make([]*models.MessageRead, 0)
	err := m.db.Select(&list, "SELECT "+
		"`message_read`.*, `message`.`content`, `message`.`type` "+
		"FROM `message_read` "+
		"INNER JOIN `message` ON (`message_read`.`message_id` = `message`.`message_id`) "+
		"WHERE  `message_read`.`app_id` = ? AND `message_read`.`user_id` = ? AND `message_read`.`status` = ? "+
		"ORDER BY `message_read`.`created_at` ASC", appId, userId, status)
	return list, err
}

// Create 创建
func (m *MessageRead) Create(model *models.MessageRead) error {
	model.CreatedAt = time.Now()
	model.UpdatedAt = model.CreatedAt
	result, err := m.db.Exec(
		"INSERT INTO `message_read` (`message_id`, `app_id`, `user_id`, `group_id`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?)",
		model.MessageId,
		model.AppId,
		model.UserId,
		model.GroupId,
		model.CreatedAt,
		model.UpdatedAt,
	)
	if err != nil {
		return err
	}

	model.Id, err = result.LastInsertId()
	return err
}

// Delete 删除数据
func (m *MessageRead) Delete(id int64) (int64, error) {
	result, err := m.db.Exec("DELETE FROM `message_read` WHERE `id` = ?", id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

// UpdateStatus 修改状态
func (m *MessageRead) UpdateStatus(id int64, status int) (int64, error) {
	result, err := m.db.Exec("UPDATE `message_read` SET `status` = ? WHERE `id` = ?", status, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
