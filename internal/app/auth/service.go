package auth

import (
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/pkg/redis"
)

type Service struct {
	redis *redis.RedisClient
}

func NewAuthService(cfg config.Config) *Service {

	return &Service{redis: redis.GetRedisClient(cfg)}
}
