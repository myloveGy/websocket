package cache

import (
	"fmt"
	"time"
	"websocket/models"
	"websocket/service"
)

type UserCache struct {
	rds *service.RedisClient
}

func NewUserCache(rds *service.RedisClient) *UserCache {
	return &UserCache{rds: rds}
}

func (u *UserCache) Get(accessToken string) (*models.User, error) {
	modelUser := &models.User{}
	if err := u.rds.Get(fmt.Sprintf("user:%s", accessToken), modelUser); err != nil {
		return nil, err
	}

	return modelUser, nil
}

func (u *UserCache) Set(accessToken string, m *models.User) error {
	return u.rds.Set(fmt.Sprintf("user:%s", accessToken), m, time.Hour)
}
