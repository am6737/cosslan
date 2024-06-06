package handler

import (
	nodev1 "github.com/cossteam/cosslan/internal/api/http/v1/node"
	"github.com/cossteam/cosslan/internal/app/dto"
	"github.com/cossteam/cosslan/internal/app/node"
	pkghttp "github.com/cossteam/cosslan/pkg/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

var _ nodev1.ServerInterface = &NodeHandler{}

type NodeHandler struct {
	useCase node.UseCase
}

func NewNodeHandler(useCase node.UseCase) *NodeHandler {
	return &NodeHandler{useCase: useCase}
}

func (n *NodeHandler) Node(ctx echo.Context) error {
	err := n.useCase.CreateNode(ctx.Request().Context(), dto.CreateNodeDTO{})
	if err != nil {
		return pkghttp.Fail(ctx, http.StatusInternalServerError, "create user error")

	}
	return pkghttp.Success(ctx, "create node", nil)
}
