package middleware

import (
	"testing"
	"websocket/utils"
)

func Test_verifyEmptyKeys(t *testing.T) {
	type args struct {
		data map[string]interface{}
		keys []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "测试正常",
			args: args{
				data: map[string]interface{}{
					"app_id": "456",
					"time":   utils.DateTime(),
					"sign":   "456",
				},
				keys: []string{"app_id", "time", "sign"},
			},
			want:  "",
			want1: false,
		},
		{
			name: "测试失败",
			args: args{
				data: map[string]interface{}{
					"time": utils.DateTime(),
					"sign": "456",
				},
				keys: []string{"app_id", "time", "sign"},
			},
			want:  "app_id",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := verifyEmptyKeys(tt.args.data, tt.args.keys)
			if got != tt.want {
				t.Errorf("verifyEmptyKeys() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("verifyEmptyKeys() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
