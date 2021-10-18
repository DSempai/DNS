package main

import (
	"DNS/config"
	"DNS/logger"
	"context"
	"expvar"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	log := logger.Initialize()
	log.Info("Starting")

	cfg := config.StartupConfig{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Info(err)
		os.Exit(1)
	}

	expvar.Publish("config_startup", expvar.Func(func() interface{} {
		return cfg
	}))

	httpServer, closer := bootstrapServer(log, cfg)
	defer closer()

	errChan := make(chan error, 1)
	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case sig := <-sigChan:
		log.Debug("Caught signal", sig)
	case err := <-errChan:
		log.Info("Error listen and serve:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	httpServer.Shutdown(ctx)
}
