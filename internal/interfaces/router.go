package interfaces

import (
	"context"
	"fmt"
	v1 "github.com/cossteam/cosslan/internal/api/http/v1"
	"github.com/cossteam/cosslan/internal/app/auth"
	"github.com/cossteam/cosslan/internal/interfaces/handler"
	pkgmiddleware "github.com/cossteam/cosslan/pkg/http/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
	"go.uber.org/zap"
)

type Router struct {
	E *echo.Echo
}

// loadSwagger 加载 Swagger 文档并处理错误
func loadSwagger(getSwaggerFunc func() (*openapi3.T, error)) (*openapi3.T, error) {
	swagger, err := getSwaggerFunc()
	if err != nil {
		return nil, fmt.Errorf("Error loading swagger spec: %w", err)
	}
	swagger.Servers = nil // 清除 servers 数组，避免验证服务器名称
	return swagger, nil
}

// NewRouter 设置 Echo 路由和应用程序的路由。
func NewRouter(
	handler *handler.Handler,
	logger *zap.Logger,
	authCase auth.UseCase,
) (*Router, error) {

	e := echo.New()

	swagger, err := loadSwagger(v1.GetSwagger)
	if err != nil {
		return nil, err
	}

	v1.RegisterHandlers(e, handler)

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

	validatorOptions := &oapimiddleware.Options{
		ErrorHandler: pkgmiddleware.HandleOpenAPIError,
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				return pkgmiddleware.HandleOpenApiAuthentication(ctx, authCase, input)
			},
		},
	}

	e.Use(oapimiddleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

	return &Router{
		E: e,
	}, nil
}
