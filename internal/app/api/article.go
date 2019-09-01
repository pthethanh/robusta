package api

import (
	"github.com/pthethanh/robusta/internal/app/article"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func newArticleHandler(policySrv article.PolicyService, es event.PublishSubscriber) (*article.Handler, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, closer, err
	}
	closer.Append(mongoCloser)

	repo := article.NewMongoRepository(s)
	conf := article.LoadConfigFromEnv()
	srv := article.NewService(conf, repo, policySrv, es)
	closer.Add(srv.Close)

	handler := article.NewHTTPHandler(srv)
	return handler, closer, nil
}
