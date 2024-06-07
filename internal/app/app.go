package app

import (
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/auth"
	"github.com/cossteam/cosslan/internal/app/line"
	"github.com/cossteam/cosslan/internal/app/node"
	"github.com/cossteam/cosslan/internal/app/user"
)

type App struct {
	UserCase user.UseCase
	NodeCase node.UseCase
	LineCase line.UseCase
	AuthCase auth.UseCase
}

func NewApp(cfg config.Config) *App {
	lineCase := line.NewLineUseCase(cfg)
	userCase := user.NewUserUseCase(cfg)
	nodeCase := node.NewNodeUseCase(cfg)
	authCase := auth.NewUseCase(cfg, userCase)
	return &App{UserCase: userCase, NodeCase: nodeCase, LineCase: lineCase, AuthCase: authCase}
}
