package api

import (
	"github.com/pthethanh/robusta/internal/app/user"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func newUserService() (*user.Service, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, closer, err
	}
	closer.Append(mongoCloser)
	repo := user.NewMongoDBRespository(s)
	return user.New(repo), closer, nil
}

func newUserHandler(srv *user.Service) *user.Handler {
	return user.NewHandler(srv)
}
