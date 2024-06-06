package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/app/line"
	"github.com/cossteam/cosslan/internal/app/node"
	"github.com/cossteam/cosslan/internal/app/user"
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

var path = ""

func init() {
	flag.StringVar(&path, "config", "config/config.yaml", "配置文件")
	flag.Parse()
}

func main() {

	// 读取配置文件
	cfg, err := config.LoadConfig(path)
	if err != nil {
		fmt.Printf("无法读取配置文件: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	logger, err := log.SetupLogger(cfg.Server.LogLevel)
	if err != nil {
		fmt.Printf("无法设置日志: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	lineCase := line.NewLineUseCase(cfg)
	userCase := user.NewUserUseCase(cfg)
	nodeCase := node.NewNodeUseCase(cfg)

	// Initialize handlers
	lineHandler := handler.NewLineHandler(lineCase)
	nodeHandler := handler.NewNodeHandler(nodeCase)
	userHandler := handler.NewUserHandler(userCase)

	// Initialize the router
	e := router.NewRouter(lineHandler, nodeHandler, userHandler, logger)

	// Run the server in a goroutine so that it doesn't block
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a signal
	<-quit
	fmt.Println("Shutting down the server...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
