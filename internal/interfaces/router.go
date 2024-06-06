package interfaces

import (
	linev1 "github.com/cossteam/cosslan/internal/api/http/v1/line"
	nodev1 "github.com/cossteam/cosslan/internal/api/http/v1/node"
	userv1 "github.com/cossteam/cosslan/internal/api/http/v1/user"
	"github.com/cossteam/cosslan/internal/interfaces/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// NewRouter sets up the Echo router with the application's routes.
func NewRouter(
	lineHandler *handler.LineHandler,
	nodeHandler *handler.NodeHandler,
	userHandler *handler.UserHandler,
	logger *zap.Logger,
) *echo.Echo {
	e := echo.New()

	userv1.RegisterHandlers(e, userHandler)
	linev1.RegisterHandlers(e, lineHandler)
	nodev1.RegisterHandlers(e, nodeHandler)
	e.Use(middleware.CORS(), middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	return e
}
