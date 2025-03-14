// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"ai-seller/config"
	v1 "ai-seller/internal/controller/http"
	"ai-seller/internal/repo/persistent"
	"ai-seller/internal/usecase/product"
	"ai-seller/pkg/httpserver"
	"ai-seller/pkg/logger"
	"ai-seller/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	// Logger
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	useCases := product.New(
		persistent.NewAuthRepo(pg),
		persistent.NewProductRepo(pg),
		persistent.NewIntegrationRepo(pg),
	)

	// HTTP Server
	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	v1.NewRouter(httpServer.Engine, l, useCases)

	httpServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
