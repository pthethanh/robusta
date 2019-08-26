package api

import (
	"github.com/pthethanh/robusta/internal/app/reaction"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func createReactionHandler(es event.Publisher) (*reaction.Handler, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, closer, err
	}
	closer.Append(mongoCloser)

	repo := reaction.NewMongoDBRepository(s)
	conf := reaction.LoadConfigFromEnv()
	srv := reaction.NewService(conf, repo, es)
	handler := reaction.NewHandler(srv)
	return handler, closer, nil
}
