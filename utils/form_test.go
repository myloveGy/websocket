package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestValidError_Error(t *testing.T) {
	type fields struct {
		Key     string
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "测试",
			fields: fields{Key: "name", Message: "error"},
			want:   "error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ValidError{
				Key:     tt.fields.Key,
				Message: tt.fields.Message,
			}
			if got := v.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidErrors_Error(t *testing.T) {
	tests := []struct {
		name string
		v    ValidErrors
		want string
	}{
		{
			name: "测试",
			v:    ValidErrors{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidErrors_Errors(t *testing.T) {
	tests := []struct {
		name string
		v    ValidErrors
		want []string
	}{
		{
			name: "测试",
			v:    ValidErrors{},
			want: nil,
		},
		{
			name: "测试二",
			v: ValidErrors{
				&ValidError{Key: "123", Message: "456"},
			},
			want: []string{"456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBindAndValid(t *testing.T) {
	type args struct {
		c *gin.Context
		v interface{}
	}

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Request, _ = http.NewRequest("POST", "/hola", strings.NewReader(`{"type":"456"}`))
	context.Request.Header.Set("Content-type", "application/json")
	//context.JSON(http.StatusNoContent, gin.H{"foo": "bar"})
	type param struct {
		Type string `json:"type" binding:"required"`
	}

	context1, _ := gin.CreateTestContext(httptest.NewRecorder())
	context1.Request, _ = http.NewRequest("POST", "/hola", strings.NewReader(`{"type":"456","str":""}`))
	context1.Request.Header.Set("Content-type", "application/json")
	type param1 struct {
		param
		S string `json:"str" binding:"required"`
	}

	tests := []struct {
		name  string
		args  args
		want  bool
		want1 ValidErrors
	}{
		{
			name: "测试用例一",
			args: args{
				c: context,
				v: &param{},
			},
			want:  false,
			want1: nil,
		},
		{
			name: "测试用例二",
			args: args{
				c: context,
				v: &param1{},
			},
			want:  true,
			want1: nil,
		},
		{
			name: "测试用例三",
			args: args{
				c: context1,
				v: &param1{},
			},
			want: true,
			want1: ValidErrors{
				&ValidError{
					Key:     "param1.S",
					Message: "Key: 'param1.S' Error:Field validation for 'S' failed on the 'required' tag",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BindAndValid(tt.args.c, tt.args.v)
			if got != tt.want {
				t.Errorf("BindAndValid() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BindAndValid() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
