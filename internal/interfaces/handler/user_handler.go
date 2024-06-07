package handler

import (
	"github.com/cossteam/cosslan/internal/app/dto"
	"github.com/cossteam/cosslan/internal/app/user"
	pkghttp "github.com/cossteam/cosslan/pkg/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	useCase user.UseCase
}

func NewUserHandler(useCase user.UseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (u *UserHandler) Login(ctx echo.Context) error {
	err := u.useCase.CreateUser(ctx.Request().Context(), dto.CreateUserDTO{})
	if err != nil {
		return pkghttp.Fail(ctx, http.StatusInternalServerError, "create user error")
	}

	return pkghttp.Success(ctx, "create user success", nil)
}
