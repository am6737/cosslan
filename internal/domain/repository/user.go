package repository

import (
	"context"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/domain/filter"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	Get(ctx context.Context, id string) (*entity.User, error)
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, filter *filter.UserFilter) ([]*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}
