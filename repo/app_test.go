package repo

import (
	"testing"

	"github.com/jinxing-go/mysql"
	"github.com/stretchr/testify/assert"
)

func newTestAppRepo(t *testing.T) *App {
	mySQL := mysql.NewTestMySQL(t, "../testdata/websocket.sql")
	return NewApp(mySQL)
}

func TestAppRepo_FindByAppId(t *testing.T) {
	appRepo := newTestAppRepo(t)

	t.Run("测试正常", func(t *testing.T) {
		_, err := appRepo.FindByAppId("2020110306161001")
		assert.NoError(t, err)
	})

	t.Run("测试失败", func(t *testing.T) {
		_, err1 := appRepo.FindByAppId("1000")
		assert.Error(t, err1)
	})
}

func TestAppRepo_FindById(t *testing.T) {
	appRepo := newTestAppRepo(t)
	t.Run("测试正常", func(t *testing.T) {
		_, err := appRepo.FindById(1)
		assert.NoError(t, err)
	})

	t.Run("测试失败", func(t *testing.T) {
		_, err1 := appRepo.FindById(100)
		assert.Error(t, err1)
	})
}
