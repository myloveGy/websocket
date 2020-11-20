package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"websocket/connection"
)

func newTestUser() *User {
	db := connection.NewDB()
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
