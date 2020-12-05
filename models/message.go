package models

import (
	"time"
)

type Message struct {
	MessageId int64     `db:"message_id" json:"message_id"`
	AppId     int64     `db:"app_id" json:"app_id"`
	Type      int       `db:"type" json:"type"`
	Content   string    `db:"content" json:"content"`
	GroupId   string    `db:"group_id" json:"group_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (*Message) TableName() string {
	return "message"
}

func (*Message) PK() string {
	return "message_id"
}
