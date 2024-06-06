package user

import (
	"context"
	"cosslan/config"
	"cosslan/internal/app/dto"
	"fmt"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	CreateUser(ctx context.Context, dto dto.CreateUserDTO) error
}

// LineUseCase implements the UseCase interface for Line.
type UserUseCase struct {
	service *Service
}

// NewLineUseCase creates a new LineUseCase.
func NewUserUseCase(cfg config.Config) UseCase {
	return &UserUseCase{service: NewService(cfg)}
}

func (u *UserUseCase) CreateUser(ctx context.Context, dto dto.CreateUserDTO) error {
	fmt.Println("CreateUser", dto)
	return nil
}
