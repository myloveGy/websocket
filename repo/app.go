package repo

import (
	"github.com/jmoiron/sqlx"
	"websocket/models"
)

type App struct {
	db *sqlx.DB
}

func NewApp(db *sqlx.DB) *App {
	return &App{db: db}
}

// 查询
func (a *App) FindByAppId(appId string) (*models.App, error) {
	app := &models.App{}
	if err := a.db.Get(app, "SELECT * FROM `app` WHERE `app_id` = ?", appId); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) FindById(id int64) (*models.App, error) {
	app := &models.App{}
	if err := a.db.Get(app, "SELECT * FROM `app` WHERE `id` = ?", id); err != nil {
		return nil, err
	}

	return app, nil
}
