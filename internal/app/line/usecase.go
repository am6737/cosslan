package line

import (
	"context"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/dto"
	"go.uber.org/zap"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	CreateLine(ctx context.Context, dto dto.CreateLineDTO) error
}

// LineUseCase implements the UseCase interface for Line.
type LineUseCase struct {
	service *Service
	logger  *zap.Logger
}

// NewLineUseCase creates a new LineUseCase.
func NewLineUseCase(cfg config.Config) *LineUseCase {
	return &LineUseCase{service: NewLineService(cfg)}
}

// CreateLine handles the use case for creating a Line.
func (uc *LineUseCase) CreateLine(ctx context.Context, dto dto.CreateLineDTO) error {
	fmt.Println("创建线路", dto)
	return nil
}
