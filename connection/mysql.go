package connection

import (
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"websocket/config"
)

func NewDB() *sqlx.DB {
	defaultName := "default"
	configValue, ok := config.App.DB[defaultName]
	if !ok {
		log.Fatalln(errors.New(defaultName + ": 数据库配置为空"))
	}

	fmt.Printf("Mysql: configValue.Driver = %s, configValue.Dsn = %s \n", configValue.Driver, configValue.Dsn)
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

	return db
}
