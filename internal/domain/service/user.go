package service

import (
	"context"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence"
)

var _ UserService = &UserServiceImpl{}

type UserService interface {
	GetUser(ctx context.Context, id string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
}

type UserServiceImpl struct {
	repo *persistence.Repositories
}

func NewUserServiceImpl(cfg config.Config) UserService {
	return &UserServiceImpl{repo: persistence.GetRepositories(cfg)}
}

func (u *UserServiceImpl) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return u.repo.UserRepo.Get(ctx, id)
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, user *entity.User) error {
	return u.repo.UserRepo.Create(ctx, user)
}
