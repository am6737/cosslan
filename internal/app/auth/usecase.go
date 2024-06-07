package auth

import (
	"context"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/user"
)

// UseCase defines the interface for Line use cases.
type UseCase interface {
	Validate(ctx context.Context, token string) error
}

type AuthUseCase struct {
	service  *Service
	userCase user.UseCase
}

// NewLineUseCase creates a new LineUseCase.
func NewUseCase(cfg config.Config, userCase user.UseCase) *AuthUseCase {
	return &AuthUseCase{service: NewAuthService(cfg), userCase: userCase}
}

func (a *AuthUseCase) Validate(ctx context.Context, token string) error {
	fmt.Println("鉴权", token)
	return nil
}
