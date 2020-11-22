package repo

import (
	"github.com/jinxing-go/mysql"
	"websocket/models"
)

type App struct {
	*mysql.MySQl
}

func NewApp(mysql *mysql.MySQl) *App {
	return &App{mysql}
}

// 查询
func (a *App) FindByAppId(appId string) (*models.App, error) {
	app := &models.App{AppId: appId}
	if err := a.Find(app); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) FindById(id int64) (*models.App, error) {
	app := &models.App{Id: id}
	if err := a.Find(app); err != nil {
		return nil, err
	}

	return app, nil
}
