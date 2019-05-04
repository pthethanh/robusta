package api

import (
	"fmt"

	"github.com/pthethanh/robusta/internal/app/article"
	"github.com/pthethanh/robusta/internal/app/db"
	"github.com/pthethanh/robusta/internal/pkg/glog"
)

func newArticleHandler(conns *InfraConns) (*article.Handler, error) {
	var articleRepo article.Repository
	switch conns.Databases.Type {
	case db.TypeMongoDB:
		articleRepo = article.NewMongoRepository(conns.Databases.MongoDB)
	default:
		return nil, fmt.Errorf("database type not supported: %s", conns.Databases.Type)
	}
	logger := glog.New().WithField("package", "article")
	srv := article.NewService(articleRepo, logger)
	handler := article.New(srv, logger)
	return handler, nil
}
