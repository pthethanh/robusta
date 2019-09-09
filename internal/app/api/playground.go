package api

import (
	"github.com/pthethanh/robusta/internal/app/playground"
	client "github.com/pthethanh/robusta/internal/pkg/playground"
)

func newPlaygroundHandler() *playground.Handler {
	runner := client.New(client.LoadConfigFromEnv())
	srv := playground.NewService(runner)
	return playground.New(srv)
}
