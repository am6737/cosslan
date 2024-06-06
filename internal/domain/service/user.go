package service

import (
	"context"
	"cosslan/config"
	"cosslan/internal/app/dto"
	"cosslan/internal/infra/persistence"
	"fmt"
)

var _ UserService = &UserServiceImpl{}

type UserService interface {
	CreateUser(ctx context.Context, dto dto.CreateUserDTO) error
}

type UserServiceImpl struct {
	repo *persistence.Repositories
}

func NewUserServiceImpl(cfg config.Config) UserService {
	return &UserServiceImpl{repo: persistence.GetRepositories(cfg)}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, dto dto.CreateUserDTO) error {
	fmt.Println("创建用户")
	return nil
}
