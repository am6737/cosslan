package node

import (
	"context"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/dto"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	CreateNode(ctx context.Context, dto dto.CreateNodeDTO) error
}

// LineUseCase implements the UseCase interface for Line.
type NodeUseCase struct {
	service *Service
}

// NewLineUseCase creates a new LineUseCase.
func NewNodeUseCase(cfg config.Config) UseCase {
	return &NodeUseCase{service: NewService(cfg)}
}

func (n *NodeUseCase) CreateNode(ctx context.Context, dto dto.CreateNodeDTO) error {
	fmt.Println("创建节点")
	return nil
}
