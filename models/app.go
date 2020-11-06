package models

import (
	"time"

	"websocket/global"
)

type App struct {
	Id        int       `db:"id"`
	AppId     string    `db:"app_id"`
	AppSecret string    `db:"app_secret"`
	AppName   string    `db:"app_name"`
	VerifyUrl string    `db:"verify_url"`
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// 查询
func (a *App) Find() error {
	return global.DB.Get(a, "SELECT * FROM `app` WHERE `app_id` = ?", a.AppId)
}
