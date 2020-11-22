package connection

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinxing-go/mysql"
	"websocket/config"
)

func NewMySQL() *mysql.MySQl {
	defaultName := "default"
	configValue, ok := config.App.DB[defaultName]
	if !ok {
		fmt.Println("MySQL connection config is empty ")
		panic(errors.New(defaultName + ": config is empty"))
	}

	return mysql.NewMySQL(configValue)
}
