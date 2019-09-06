package api

import (
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"

	"github.com/pthethanh/robusta/internal/app/tutorial"
)

func newTutorialHandler(policySrv tutorial.PolicyService) (*tutorial.Handler, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, closer, err
	}
	closer.Append(mongoCloser)
	repo := tutorial.NewMongoRepository(s)
	srv := tutorial.NewService(repo, policySrv)
	handler := tutorial.NewHTTPHandler(srv)
	return handler, closer, nil
}
