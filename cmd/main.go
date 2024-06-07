package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app"
	router "github.com/cossteam/cosslan/internal/interfaces"
	"github.com/cossteam/cosslan/internal/interfaces/handler"
	"github.com/cossteam/cosslan/pkg/log"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "config/config.yaml", "配置文件")
	flag.Parse()
}

func main() {
	//加载日志与配置文件
	cfg, logger := setup()

	//加载app
	ap := app.NewApp(cfg)

	//初始化路由
	handlers := initializeHandlers(ap)
	r, err := router.NewRouter(handlers, logger, ap.AuthCase)
	if err != nil {
		logger.Fatal("Failed to initialize router", zap.Error(err))
	}

	//启动服务
	startServer(r, logger)
}

func setup() (config.Config, *zap.Logger) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("无法读取配置文件: %v\n", err)
		os.Exit(1)
	}

	logger, err := log.SetupLogger(cfg.Server.LogLevel)
	if err != nil {
		fmt.Printf("无法设置日志: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	return cfg, logger
}

func initializeHandlers(ap *app.App) *handler.Handler {
	lineHandler := handler.NewLineHandler(ap.LineCase)
	nodeHandler := handler.NewNodeHandler(ap.NodeCase)
	userHandler := handler.NewUserHandler(ap.UserCase)

	return handler.NewHandler(userHandler, nodeHandler, lineHandler)
}

func startServer(r *router.Router, logger *zap.Logger) {
	go func() {
		if err := r.E.Start(":8080"); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := r.E.Shutdown(ctx); err != nil {
		r.E.Logger.Fatal(err)
	}
}
