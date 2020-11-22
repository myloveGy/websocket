package repo

import (
	"github.com/jinxing-go/mysql"
)

type Message struct {
	*mysql.MySQl
}

func NewMessage(mysql *mysql.MySQl) *Message {
	return &Message{mysql}
}
