package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
	"websocket/global"
)

func TestMessage_Create(t *testing.T) {
	global.NewConnect("default")

	type args struct {
		db *sqlx.DB
	}

	tests := []struct {
		name    string
		fields  *Message
		args    args
		wantErr bool
	}{
		{
			name: "测试新增成功",
			fields: &Message{
				AppId:     1,
				Type:      1,
				Content:   `{"username":"jinxing.liu","age":28}`,
				CreatedAt: time.Now(),
			},
			args: args{
				db: global.DB,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.Create(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			fmt.Printf("创建数据成功: %v", tt.fields)

			// 需要删除数据
			if _, err := tt.fields.Delete(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
