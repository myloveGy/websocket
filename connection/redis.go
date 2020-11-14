package connection

import (
	"github.com/go-redis/redis/v8"
	"websocket/config"
)

var RedisDB *redis.Client

func NewRedisDb(name ...string) {
	defaultName := "default"
	if name != nil && name[0] != "" {
		defaultName = name[0]
	}

	configRedis, ok := config.App.Redis[defaultName]
	if !ok {
		panic("init redis error: config is empty")
	}

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     configRedis.Addr,
		Password: configRedis.Password,
		DB:       configRedis.DB,
	})
}
