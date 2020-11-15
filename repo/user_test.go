package repo

import (
	"testing"
	"websocket/connection"
)

func TestUserRepo_FindByUsername(t *testing.T) {
	db := connection.NewDB()
	user := NewUser(db)
	type args struct {
		username string
	}

	tests := []struct {
		name    string
		fields  *User
		args    args
		wantErr bool
	}{
		{
			name:    "查询正常",
			fields:  user,
			args:    args{username: "jinxing.liu"},
			wantErr: false,
		},
		{
			name:    "查询错误",
			fields:  user,
			args:    args{username: "haha.ha"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.fields.FindByUsername(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("FindByUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
