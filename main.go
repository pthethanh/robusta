package main

import (
	"flag"

	"github.com/pthethanh/robusta/internal/app/api"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/http/server"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func main() {
	env := flag.String("env", "", "env file")
	flag.Parse()
	if *env != "" {
		if err := envconfig.SetEnvFromFile(*env); err != nil {
			log.Panicf("failed to set env, err: %v", err)
		}
	}
	log.Init(log.Fields{
		"service": "robusta",
	})
	log.Infof("initializing HTTP routing...")
	router, closer, err := api.NewRouter()
	if err != nil {
		log.Panicf("failed to init routing, err: %v", err)
	}
	defer closer.Close()
	serverConf := server.LoadConfigFromEnv()
	server.ListenAndServe(serverConf, router)
}
