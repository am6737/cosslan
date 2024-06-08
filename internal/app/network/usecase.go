package node

import (
	"context"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/dto"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	CreateNetwork(ctx context.Context, dto dto.CreateNetworkDTO) error
}

// LineUseCase implements the UseCase interface for Line.
type NodeUseCase struct {
	service *Service
}

// NewLineUseCase creates a new LineUseCase.
func NewNetworkUseCase(cfg config.Config) UseCase {
	return &NodeUseCase{service: NewService(cfg)}
}

func (n *NodeUseCase) CreateNetwork(ctx context.Context, dto dto.CreateNetworkDTO) error {
	err := n.service.svc.CreateNetwork(ctx, dto)
	if err != nil {
		return err
	}
	return nil
}
