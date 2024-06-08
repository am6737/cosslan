package repository

import (
	"context"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/domain/filter"
)

type NetworkRepository interface {
	Update(ctx context.Context, network *entity.Network) error
	Get(ctx context.Context, id string)
	Create(ctx context.Context, network *entity.Network) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, filter *filter.NetworkFilter) ([]*entity.Network, error)
}
