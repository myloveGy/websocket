package utils

import (
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "日期",
			want: time.Now().Format(DateLayout),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Date(); got != tt.want {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "日期时间",
			want: time.Now().Format(DateTimeLayout),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateTime(); got != tt.want {
				t.Errorf("DateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
