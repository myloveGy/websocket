package repo

import (
	"fmt"
	"testing"
	"websocket/connection"
	"websocket/models"
)

func newTestMessageRead() *MessageRead {
	db := connection.NewDB()
	return NewMessageRead(db)
}

func TestMessageRead_Create(t *testing.T) {
	messageReadRepo := newTestMessageRead()
	tests := []struct {
		name    string
		fields  *MessageRead
		args    *models.MessageRead
		wantErr bool
	}{
		{
			name:   "测试新增成功",
			fields: messageReadRepo,
			args: &models.MessageRead{
				AppId:     1,
				MessageId: 1,
				UserId:    "1",
				GroupId:   "",
			},
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
			if _, err := tt.fields.Delete(tt.args.Id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMessageRead_FindAll(t *testing.T) {
	messageReadRepo := newTestMessageRead()
	type args struct {
		AppId  int64
		UserId string
		Status int
	}

	tests := []struct {
		name    string
		fields  *MessageRead
		args    *args
		want    []*models.MessageRead
		wantErr bool
	}{
		{
			name:    "测试用例",
			args:    &args{AppId: 1, UserId: "1", Status: 1},
			fields:  messageReadRepo,
			want:    []*models.MessageRead{{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.FindAll(tt.args.AppId, tt.args.UserId, tt.args.Status)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageRead_UpdateStatus(t *testing.T) {
	messageReadRepo := newTestMessageRead()
	type args struct {
		Status int
		Id     int64
	}
	tests := []struct {
		name    string
		fields  *MessageRead
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "测试正确",
			args:    args{Id: 1, Status: 2},
			fields:  messageReadRepo,
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.UpdateStatus(tt.args.Id, tt.args.Status)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
