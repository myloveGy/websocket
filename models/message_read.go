package models

import (
	"time"
)

type MessageRead struct {
	Id        int64     `db:"id" json:"id"`
	MessageId int64     `db:"message_id" json:"message_id"`
	AppId     int64     `db:"app_id" json:"app_id"`
	UserId    string    `db:"user_id" json:"user_id"`
	GroupId   string    `db:"group_id" json:"group_id"`
	Status    int       `db:"status" json:"status"`
	Content   string    `db:"content" json:"content"`
	Type      int       `db:"type" json:"type"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (*MessageRead) TableName() string {
	return "message_read"
}

func (*MessageRead) PK() string {
	return "id"
}
