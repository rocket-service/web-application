package app

import (
	"context"
	"rocket-web/internal/config"
	"rocket-web/internal/router"
	"rocket-web/internal/storage/postgres"

	"go.uber.org/zap"
)

type App struct {
	router *router.Router
	log    *zap.SugaredLogger
}

func New(cfg *config.Config, log *zap.SugaredLogger) *App {
	postgres, err := postgres.New(context.Background(), log, &cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}

	err = postgres.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	router := router.New(cfg, log)
	return &App{router: router, log: log}
}

func (a *App) Run(port int32) error {
	return a.router.MustRun(port)
}

func (a *App) Close() error {
	return a.router.Close()
}
