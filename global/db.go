package global

import (
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	config "websocket/config"
)

var DB *sqlx.DB

func NewConnect(defaultConfig string) {

	configValue, ok := config.App.DB[defaultConfig]
	if !ok {
		log.Fatalln(errors.New(defaultConfig + ": 数据库配置为空"))
	}

	fmt.Printf("configValue.Driver = %s, configValue.Dsn = %s", configValue.Driver, configValue.Dsn)
	db, err := sqlx.Connect(configValue.Driver, configValue.Dsn)
	if err != nil {
		fmt.Println("error: ", err)
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(configValue.MaxIdleConns)
	db.SetMaxOpenConns(configValue.MaxOpenConns)
	if configValue.MaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(configValue.MaxLifetime) * time.Second)
	}

	DB = db
}