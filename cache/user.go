package cache

import (
	"fmt"
	"time"
	"websocket/config"
	"websocket/connection"
	"websocket/models"
)

type UserCache struct {
	rds *connection.Redis
}

func NewUserCache(rds *connection.Redis) *UserCache {
	return &UserCache{rds: rds}
}

func (u *UserCache) key(s string) string {
	return fmt.Sprintf("user:%s", s)
}

func (u *UserCache) Get(accessToken string) (*models.User, error) {
	modelUser := &models.User{}
	if err := u.rds.Get(u.key(accessToken), modelUser); err != nil {
		return nil, err
	}

	return modelUser, nil
}

func (u *UserCache) Set(accessToken string, m *models.User) error {
	return u.rds.Set(u.key(accessToken), m, time.Hour*config.App.LoginTime)
}

func (u *UserCache) Delete(accessToken string) error {
	return u.rds.Delete(u.key(accessToken))
}
