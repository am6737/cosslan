package service

import (
	"context"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/dto"
	"github.com/cossteam/cosslan/internal/infra/persistence"
)

var _ LineService = &LineServiceImpl{}

type LineService interface {
	CreateLine(ctx context.Context, dto dto.CreateLineDTO) error
}

type LineServiceImpl struct {
	repo *persistence.Repositories
}

func NewLineService(cfg config.Config) *LineServiceImpl {
	return &LineServiceImpl{repo: persistence.GetRepositories(cfg)}
}

func (l *LineServiceImpl) CreateLine(ctx context.Context, dto dto.CreateLineDTO) error {
	fmt.Println("创建线路")
	return nil
}
