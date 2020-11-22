package repo

import (
	"fmt"
	"testing"
	"websocket/models"

	"github.com/stretchr/testify/assert"

	"websocket/connection"
)

func newTestUser() *User {
	db := connection.NewMySQL()
	return NewUser(db)
}

func TestUserRepo_FindByUsername(t *testing.T) {
	user := newTestUser()
	t.Run("查询正常", func(t *testing.T) {
		u, err := user.FindByUsername("jinxing.liu")
		assert.NoError(t, err)
		assert.Equal(t, "jinxing.liu", u.Username)
	})

	t.Run("查询失败", func(t *testing.T) {
		u, err := user.FindByUsername("jinxing.liuxxx")
		assert.Error(t, err)
		assert.Nil(t, u)
	})
}

func TestUser_UpdateAccessToken(t *testing.T) {
	user := newTestUser()
	t.Run("测试修改", func(t *testing.T) {
		u, err := user.UpdateAccessToken(1, "jinxing.liu")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), u)
	})
}

func TestUser_Create(t *testing.T) {
	user := newTestUser()
	t.Run("测试失败", func(t *testing.T) {
		err := user.Create(&models.User{
			Username:    "jinxing.liu",
			Password:    "123456",
			AccessToken: "789123",
		})

		fmt.Println(err)
		assert.Error(t, err)
	})

	t.Run("测试正常", func(t *testing.T) {
		user.Exec("DELETE FROM `user` WHERE `username` = ?", "jinxing.liu1")
		err := user.Create(&models.User{
			Username:    "jinxing.liu1",
			Password:    "123456",
			Phone:       "12345678901",
			AccessToken: "789123",
		})

		fmt.Println(err)
		assert.NoError(t, err)
	})
}

func TestUser_Delete(t *testing.T) {
	row, err := newTestUser().Delete(&models.User{UserId: 100})
	assert.NoError(t, err)
	assert.Equal(t, int64(0), row)
}
