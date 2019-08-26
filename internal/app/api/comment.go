package api

import (
	"github.com/pthethanh/robusta/internal/app/comment"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func newCommentHandler(policySrv comment.PolicyService, es event.PublishSubscriber) (*comment.Handler, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, closer, err
	}
	closer.Append(mongoCloser)
	repo := comment.NewMongoRepository(s)
	conf := comment.LoadConfigFromEnv()
	srv := comment.NewService(conf, repo, policySrv, es)
	closer.Add(srv.Close)

	handler := comment.NewHandler(srv)
	return handler, closer, nil
}
