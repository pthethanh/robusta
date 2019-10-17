package api

import (
	"github.com/pthethanh/robusta/internal/app/user"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func newUserService(policy user.PolicyService, es event.Publisher) (*user.Service, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, closer, err
	}
	closer.Append(mongoCloser)
	repo := user.NewMongoDBRespository(s)
	conf := user.LoadConfigFromEnv()
	return user.New(conf, repo, policy, newJWTSignVerifier(), es), closer, nil
}

func newUserHandler(srv *user.Service) *user.Handler {
	return user.NewHandler(srv)
}
