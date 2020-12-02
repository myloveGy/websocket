package testdata

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jinxing-go/mysql"
)

func Schema() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "example.sql")
}

func NewTestMySQL(t *testing.T, filename ...string) *mysql.MySQl {
	return mysql.NewTestMySQL(t, Schema(), filename...)
}
