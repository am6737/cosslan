package handler

import v1 "github.com/cossteam/cosslan/internal/api/http/v1"

var _ v1.ServerInterface = &Handler{}

type Handler struct {
	*UserHandler
	*NodeHandler
	*LineHandler
}

func NewHandler(userHandler *UserHandler, nodeHandler *NodeHandler, lineHandler *LineHandler) *Handler {
	return &Handler{UserHandler: userHandler, NodeHandler: nodeHandler, LineHandler: lineHandler}
}
