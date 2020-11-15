package repo

import (
	"fmt"
	"testing"
	"time"
	"websocket/connection"
	"websocket/models"
)

func newTestMessageRepo() *Message {
	db := connection.NewDB()
	return NewMessage(db)
}

func TestMessage_Create(t *testing.T) {
	messageRepo := newTestMessageRepo()
	tests := []struct {
		name    string
		fields  *Message
		args    *models.Message
		wantErr bool
	}{
		{
			name: "测试新增成功",
			args: &models.Message{
				AppId:     1,
				Type:      1,
				Content:   `{"username":"jinxing.liu","age":28}`,
				CreatedAt: time.Now(),
			},
			fields:  messageRepo,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.Create(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			fmt.Printf("创建数据成功: %v \n", tt.fields)

			// 需要删除数据
			if _, err := tt.fields.Delete(tt.args.MessageId); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMessage_Delete(t *testing.T) {
	messageRepo := newTestMessageRepo()
	tests := []struct {
		name    string
		fields  *Message
		args    int64
		want    int64
		wantErr bool
	}{
		{
			name:    "测试正常",
			fields:  messageRepo,
			args:    10000,
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.Delete(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}
