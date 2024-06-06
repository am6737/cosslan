package repository

import (
	"context"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/domain/filter"
)

type LineRepository interface {
	Create(ctx context.Context, line *entity.Line) error
	Get(ctx context.Context, id string) (*entity.Line, error)
	Update(ctx context.Context, line *entity.Line) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, filter *filter.LineFilter) ([]*entity.Line, error)
}
