package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"awesomeProject3/pkg/config"
)

type Redis struct {
	client *redis.Client
}

func RedisManager() *Redis {
	return &Redis{client: config.CreateRedisDatabase()}
}

type RedisMethods interface {
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	Get(string) *redis.StringCmd
}

func (cache *Redis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return cache.client.Set(context.Background(), key, value, expiration)
}
func (cache *Redis) Get(key string) *redis.StringCmd {
	return cache.client.Get(context.Background(), key)
}
