package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pthethanh/robusta/internal/app/api"
	"github.com/pthethanh/robusta/internal/app/db"
	"github.com/pthethanh/robusta/internal/app/db/mongodb"
	"github.com/pthethanh/robusta/internal/pkg/config/env"
	"github.com/pthethanh/robusta/internal/pkg/glog"
	"github.com/pthethanh/robusta/internal/pkg/health"
)

type (
	srvConfig struct {
		DB struct {
			Type    string `envconfig:"DB_TYPE" default:"mongodb"`
			MongoDB mongodb.Config
		}
		HTTP struct {
			Address           string        `envconfig:"HTTP_ADDRESS" default:""`
			Port              int           `envconfig:"PORT" default:"8080"`
			ReadTimeout       time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"5m"`
			WriteTimeout      time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"5m"`
			ReadHeaderTimeout time.Duration `envconfig:"HTTP_READ_HEADER_TIMEOUT" default:"30s"`
			ShutdownTimeout   time.Duration `envconfig:"HTTP_SHUTDOWN_TIMEOUT" default:"10s"`
		}
	}
)

func main() {
	var conf srvConfig
	envconfig.Load(&conf)
	logger := glog.New()

	conns := initInfraConns(&conf, logger)
	defer conns.Close()

	logger.Infof("initializing HTTP routing...")
	router, err := api.Init(conns)
	if err != nil {
		logger.Panicf("failed to init routing, err: %v", err)
	}
	addr := fmt.Sprintf("%s:%d", conf.HTTP.Address, conf.HTTP.Port)
	httpServer := http.Server{
		Addr:              addr,
		Handler:           router,
		ReadTimeout:       conf.HTTP.ReadTimeout,
		WriteTimeout:      conf.HTTP.WriteTimeout,
		ReadHeaderTimeout: conf.HTTP.ReadHeaderTimeout,
	}

	logger.Infof("starting HTTP server...")
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Panicf("http.ListenAndServe() error: %v", err)
		}
	}()

	// tell the world that we're ready
	health.Ready()
	logger.Infof("HTTP Server is listening at: %v", addr)

	// gracefully shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	<-signals
	ctx, cancel := context.WithTimeout(context.Background(), conf.HTTP.ShutdownTimeout)
	defer cancel()
	logger.Infof("shutting down http server...")
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Errorf("http server shutdown with error: %v", err)
	}
	// shutdown background services goes here
}

// initInfraConns init underlying infrastructure connections
func initInfraConns(conf *srvConfig, l glog.Logger) *api.InfraConns {
	conns := &api.InfraConns{}
	conns.Databases.Type = conf.DB.Type

	switch conf.DB.Type {
	case db.TypeMongoDB:
		s, err := mongodb.Dial(&conf.DB.MongoDB, l)
		if err != nil {
			l.Panicf("failed to dial to target server, err: %v", err)
		}
		conns.Databases.MongoDB = s
	}
	return conns
}
