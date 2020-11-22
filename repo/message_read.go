package repo

import (
	"github.com/jinxing-go/mysql"
	"websocket/models"
)

type MessageRead struct {
	*mysql.MySQl
}

func NewMessageRead(mysql *mysql.MySQl) *MessageRead {
	return &MessageRead{mysql}
}

func (m *MessageRead) FindAll(appId int64, userId string, status int) ([]*models.MessageRead, error) {
	list := make([]*models.MessageRead, 0)
	err := m.DB.Select(&list, "SELECT "+
		"`message_read`.*, `message`.`content`, `message`.`type` "+
		"FROM `message_read` "+
		"INNER JOIN `message` ON (`message_read`.`message_id` = `message`.`message_id`) "+
		"WHERE  `message_read`.`app_id` = ? AND `message_read`.`user_id` = ? AND `message_read`.`status` = ? "+
		"ORDER BY `message_read`.`created_at` ASC", appId, userId, status)
	return list, err
}

// UpdateStatus 修改状态
func (m *MessageRead) UpdateStatus(id int64, status int) (int64, error) {
	return m.Update(&models.MessageRead{
		Id:     id,
		Status: status,
	}, "status")
}
