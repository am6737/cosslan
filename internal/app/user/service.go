package user

import (
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/domain/service"
)

type Service struct {
	svc service.UserService
}

func NewService(cfg config.Config) *Service {
	return &Service{svc: service.NewUserServiceImpl(cfg)}
}
