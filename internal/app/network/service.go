package node

import (
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/domain/service"
)

type Service struct {
	svc service.NetworkService
}

func NewService(cfg config.Config) *Service {
	return &Service{svc: service.NewNetworkService(cfg)}
}
