package service

import (
	"context"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/dto"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence"
)

type NetworkService interface {
	CreateNetwork(ctx context.Context, dto dto.CreateNetworkDTO) error
}

type NetworkServiceImpl struct {
	repo *persistence.Repositories
}

func NewNetworkService(cfg config.Config) *NetworkServiceImpl {
	return &NetworkServiceImpl{
		repo: persistence.GetRepositories(cfg),
	}
}

func (n *NetworkServiceImpl) CreateNetwork(ctx context.Context, dto dto.CreateNetworkDTO) error {
	err := n.repo.NetworkRepo.Create(ctx, &entity.Network{
		Address:    dto.Address,
		SubnetMask: dto.SubnetMask,
		UserID:     dto.UserID,
	})
	if err != nil {
		return err
	}
	return nil
}
