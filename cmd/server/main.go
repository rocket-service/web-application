package main

import (
	"os"
	"os/signal"
	"rocket-web/internal/app"
	"rocket-web/internal/config"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Infof("starting server in %s mode", cfg.Env)

	app := app.New(cfg, log)

	err := app.Run(cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Info("stopping server")
	app.Close()
}

func setupLogger(env string) *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	switch env {
	case config.EnvLocal:
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

	case config.EnvProduction:
		logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}

	sugaredLogger := logger.Sugar()

	return sugaredLogger
}
