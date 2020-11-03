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
	Status    int       `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (a *App) Find() error {
	return global.DB.Get(a, "SELECT * FROM `app` WHERE `app_id` = ?", a.AppId)
}
