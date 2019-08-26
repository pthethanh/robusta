package api

import (
	"github.com/pthethanh/robusta/internal/app/playground"
)

func newPlaygroundHandler() *playground.Handler {
	srv := playground.NewService()
	return playground.New(srv)
}
