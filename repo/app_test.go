package repo

import (
	"testing"
	"websocket/connection"
)

func newTestAppRepo() *App {
	db := connection.NewDB()
	return NewApp(db)
}

func TestAppRepo_FindByAppId(t *testing.T) {
	appRepo := newTestAppRepo()
	tests := []struct {
		name    string
		fields  *App
		args    string
		wantErr bool
	}{
		{
			name:    "查询正常",
			fields:  appRepo,
			args:    "2020110306161001",
			wantErr: false,
		},
		{
			name:    "查询没有数据",
			fields:  appRepo,
			args:    "100",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.fields.FindByAppId(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("FindByAppId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppRepo_FindById(t *testing.T) {
	appRepo := newTestAppRepo()
	tests := []struct {
		name    string
		fields  *App
		args    int64
		wantErr bool
	}{
		{
			name:    "查询正常",
			fields:  appRepo,
			args:    1,
			wantErr: false,
		},
		{
			name:    "查询没有数据",
			fields:  appRepo,
			args:    100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.fields.FindById(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
