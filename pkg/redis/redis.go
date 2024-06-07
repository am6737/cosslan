package redis

import (
	"context"
	"github.com/cossteam/cosslan/config"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	redisClient *RedisClient
	once        sync.Once //Once确保某些操作只被执行一次
)

type RedisClient struct {
	Client *redis.Client
}

func newRedisClient(address string, password string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	return &RedisClient{Client: client}
}

func GetRedisClient(cfg config.Config) *RedisClient {
	once.Do(func() {
		redisClient = newRedisClient(cfg.Redis.Address, cfg.Redis.Password)
	})
	return redisClient
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}
