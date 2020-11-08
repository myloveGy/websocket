package utils

import (
	"reflect"
	"testing"
)

func TestMapToString(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试用例",
			args: args{data: map[string]interface{}{
				"age":      "123",
				"username": "456",
				"b":        456,
				"my_name":  "jinxing.liu",
			}},
			want: "age=123&b=456&my_name=jinxing.liu&username=456",
		},
		{
			name: "测试用例",
			args: args{data: map[string]interface{}{
				"a":       int(1),
				"b":       int8(8),
				"c":       int32(32),
				"de":      int64(1),
				"bb":      false,
				"d":       float32(2.00),
				"e":       float64(5.00),
				"my_name": "jinxing.liu",
				"sign":    "abcde",
			}},
			want: "a=1&b=8&bb=false&c=32&d=2&de=1&e=5&my_name=jinxing.liu",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToString(tt.args.data); got != tt.want {
				t.Errorf("MapToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapStringSort(t *testing.T) {
	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "测试用例",
			args: args{data: map[string]interface{}{
				"b": "123",
				"a": "456",
				"d": "4678",
				"m": "123",
			}},
			want: []string{"a", "b", "d", "m"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMapStringSort(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMapStringSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSign(t *testing.T) {
	type args struct {
		data   map[string]interface{}
		Secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试加密",
			args: args{
				data: map[string]interface{}{
					"username": "456",
					"age":      "789",
				},
				Secret: "456",
			},
			want: "11a03d0d4ab530e0c5b406b05d9af076",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sign(tt.args.data, tt.args.Secret); got != tt.want {
				t.Errorf("Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyEmptyKeys(t *testing.T) {
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
			name: "测试一",
			args: args{data: map[string]interface{}{
				"username": "age",
				"age":      20,
				"time":     "2020-02-20 11:00:11",
			}, keys: []string{"username", "age", "time"}},
			want:  "",
			want1: false,
		},
		{
			name: "测试二",
			args: args{data: map[string]interface{}{
				"username": "age",
				"time":     "2020-02-20 11:00:11",
			}, keys: []string{"username", "age", "time"}},
			want:  "age",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := VerifyEmptyKeys(tt.args.data, tt.args.keys)
			if got != tt.want {
				t.Errorf("VerifyEmptyKeys() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("VerifyEmptyKeys() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
