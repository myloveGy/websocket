package models

import (
	"github.com/jinxing-go/mysql"
	"time"
)

type User struct {
	UserId      int64      `db:"user_id" json:"user_id"`
	Username    string     `db:"username" json:"username"`
	Phone       string     `db:"phone" json:"phone"`
	Password    string     `db:"password" json:"password"`
	Status      int        `db:"status" json:"status"`
	AccessToken string     `db:"access_token" json:"access_token"`
	CreatedAt   mysql.Time `db:"created_at" json:"created_at"`
	UpdatedAt   mysql.Time `db:"updated_at" json:"updated_at"`
}

func (*User) TableName() string {
	return "user"
}

func (*User) PK() string {
	return "user_id"
}

func (*User) TimestampsValue() interface{} {
	return mysql.Time(time.Now())
}
