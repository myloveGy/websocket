package utils

import (
	"fmt"
	"testing"
)

func TestGetHTTP(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "测试正常",
			args: args{
				message: "你好",
			},
			want: &Result{
				Code:    0,
				Content: "",
			},
			wantErr: false,
		},
		{
			name: "测试错误数据",
			args: args{
				message: "",
			},
			want: &Result{
				Code:    1,
				Content: "未获取到相关信息",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHTTP(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHTTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil || got.Code != tt.want.Code {
				t.Errorf("GetHTTP() got = %v, want %v", got, tt.want)
			}

			fmt.Printf("result: %v \n", got)
		})
	}
}
