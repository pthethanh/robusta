package api

import (
	"github.com/pthethanh/robusta/internal/app/policy"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
)

func newPolicyService() (*policy.Service, error) {
	var conf policy.CasbinConfig
	envconfig.LoadWithPrefix("CASBIN", &conf)
	enforcer := policy.NewMongoDBCasbinEnforcer(conf)
	return policy.New(enforcer)
}

func newPolicyHandler(srv *policy.Service) *policy.Handler {
	return policy.NewHandler(srv)
}
