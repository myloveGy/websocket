package models

import (
	"testing"
	"websocket/global"
)

func TestApp_Find(t *testing.T) {
	global.NewConnect("default")
	tests := []struct {
		name    string
		fields  *App
		wantErr bool
	}{
		{
			name:    "查询没有",
			fields:  &App{AppId: "123"},
			wantErr: true,
		},
		{
			name:    "查询有数据",
			fields:  &App{AppId: "2020110306161001"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.Find(); (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
