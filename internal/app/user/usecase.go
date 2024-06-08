package user

import (
	"context"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/dto"
	"github.com/cossteam/cosslan/internal/domain/entity"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	Login(ctx context.Context, dto dto.LoginUserDTO) (string, error)
	Register(ctx context.Context, dto dto.RegisterUserDTO) (*entity.User, error)
}

// LineUseCase implements the UseCase interface for Line.
type UserUseCase struct {
	service *Service
}

// NewLineUseCase creates a new LineUseCase.
func NewUserUseCase(cfg config.Config) UseCase {
	return &UserUseCase{service: NewService(cfg)}
}

func (u *UserUseCase) Login(ctx context.Context, dto dto.LoginUserDTO) (string, error) {
	//user, err := u.service.svc.GetUser(ctx, dto.Email)
	//if err != nil {
	//	return "", err
	//}
	fmt.Printf("user: %v", dto)
	return "", nil
}

func (u *UserUseCase) Register(ctx context.Context, dto dto.RegisterUserDTO) (*entity.User, error) {
	fmt.Printf("user: %v", dto)
	return nil, nil
}
