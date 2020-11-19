package connection

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	"websocket/config"
)

var ctx = context.Background()

type Redis struct {
	Prefix string
	Client *redis.Client
}

func NewRedis() *Redis {
	defaultName := "default"
	configRedis, ok := config.App.Redis[defaultName]
	if !ok {
		panic("init redis error: config is empty")
	}

	return &Redis{
		Prefix: configRedis.Prefix,
		Client: redis.NewClient(&redis.Options{
			Addr:     configRedis.Addr,
			Password: configRedis.Password,
			DB:       configRedis.DB,
		}),
	}
}

func (r *Redis) Get(name string, result interface{}) error {
	val, err := r.Client.Get(ctx, r.Prefix+name).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), result)
}

func (r *Redis) Set(name string, result interface{}, expiration time.Duration) error {

	val, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return r.Client.Set(ctx, r.Prefix+name, val, expiration).Err()
}

func (r *Redis) Delete(name string) error {
	return r.Client.Del(ctx, r.Prefix+name).Err()
}
