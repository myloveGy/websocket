package repo

import (
	"database/sql"
	"errors"
	"github.com/jinxing-go/mysql"
	"websocket/models"
)

type User struct {
	*mysql.MySQl
}

func NewUser(mysql *mysql.MySQl) *User {
	return &User{MySQl: mysql}
}

func (u *User) FindByUsername(username string) (*models.User, error) {
	user := &models.User{Username: username}
	if err := u.Find(user, "username"); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) FindByPhone(phone string) (*models.User, error) {
	user := &models.User{Phone: phone}
	if err := u.Find(user, "phone"); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) ExistsPhone(phone string, userId int64) bool {
	user := &models.User{}
	if err := u.Builder(user).
		Where("phone", phone).
		Where("user_id", "!=", userId).One(); errors.Is(err, sql.ErrNoRows) {
		return false
	}

	return true
}

func (u *User) ExistsUsername(username string, userId int64) bool {
	user := &models.User{}
	if err := u.Builder(user).
		Where("username", username).
		Where("user_id", "!=", userId).One(); errors.Is(err, sql.ErrNoRows) {
		return false
	}

	return true
}

func (u *User) UpdateAccessToken(userId int64, accessToken string) (int64, error) {
	return u.Update(&models.User{
		UserId:      userId,
		AccessToken: accessToken,
	})
}
