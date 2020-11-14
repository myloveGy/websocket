package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisClient struct {
	Prefix string `json:"prefix"` // 前缀
	Client *redis.Client
}

var ctx = context.Background()

func NewRedisClient(client *redis.Client) *RedisClient {
	return &RedisClient{
		Prefix: "ws:",
		Client: client,
	}
}

func (r *RedisClient) Get(name string, result interface{}) error {
	val, err := r.Client.Get(ctx, r.Prefix+name).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), result)
}

func (r *RedisClient) Set(name string, result interface{}, expiration time.Duration) error {

	val, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return r.Client.Set(ctx, r.Prefix+name, val, expiration).Err()
}
