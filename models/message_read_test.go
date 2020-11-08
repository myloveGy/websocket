package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
	"websocket/global"
)

func TestMessageRead_Create(t *testing.T) {
	global.NewConnect("default")

	type args struct {
		db *sqlx.DB
	}

	tests := []struct {
		name    string
		fields  *MessageRead
		args    args
		wantErr bool
	}{
		{
			name: "测试新增成功",
			fields: &MessageRead{
				AppId:     1,
				MessageId: 1,
				UserId:    "1",
				GroupId:   "",
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

func TestMessageRead_FindAll(t *testing.T) {
	global.NewConnect("default")
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  *MessageRead
		args    args
		want    []*MessageRead
		wantErr bool
	}{
		{
			name:    "测试用例",
			fields:  &MessageRead{AppId: 1, UserId: "1", Status: 1},
			args:    args{db: global.DB},
			want:    []*MessageRead{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.FindAll(tt.args.db)
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
	type args struct {
		db *sqlx.DB
	}

	global.NewConnect("default")
	tests := []struct {
		name    string
		fields  *MessageRead
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "测试正确",
			fields:  &MessageRead{Id: 1, Status: 2},
			args:    args{db: global.DB},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.UpdateStatus(tt.args.db)
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
