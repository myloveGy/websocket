package models

import (
	"time"
)

type User struct {
	UserId      int64     `db:"user_id" json:"user_id"`
	AppId       int64     `db:"app_id" json:"app_id"`
	Username    string    `db:"username" json:"username"`
	Phone       string    `db:"phone" json:"phone"`
	Password    string    `db:"password" json:"password"`
	Status      int       `db:"status" json:"status"`
	AccessToken string    `db:"access_token" json:"access_token"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}
