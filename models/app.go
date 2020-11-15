package models

import (
	"time"
)

type App struct {
	Id        int64     `db:"id" json:"id"`
	AppId     string    `db:"app_id" json:"app_id"`
	AppSecret string    `db:"app_secret" json:"app_secret"`
	AppName   string    `db:"app_name" json:"app_name"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
