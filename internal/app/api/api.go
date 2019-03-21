package api

import (
	"fmt"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/api/handler/article"
	"github.com/pthethanh/robusta/internal/app/api/handler/index"
	"github.com/pthethanh/robusta/internal/app/article"
	"github.com/pthethanh/robusta/internal/app/db"
	"github.com/pthethanh/robusta/internal/pkg/glog"
	"github.com/pthethanh/robusta/internal/pkg/health"
	"github.com/pthethanh/robusta/internal/pkg/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type (
	// InfraConns holds infrastructure services connections like MongoDB, Redis, Kafka,...
	InfraConns struct {
		Databases db.Connections
	}

	middlewareFunc = func(http.HandlerFunc) http.HandlerFunc
	route          struct {
		path        string
		method      string
		handler     http.HandlerFunc
		middlewares []middlewareFunc
	}
)

const (
	get    = http.MethodGet
	post   = http.MethodPost
	put    = http.MethodPut
	delete = http.MethodDelete
)

// Init init all handlers
func Init(conns *InfraConns) (http.Handler, error) {
	logger := glog.New()

	var articleRepo article.Repository
	switch conns.Databases.Type {
	case db.TypeMongoDB:
		articleRepo = article.NewMongoRepository(conns.Databases.MongoDB)
	default:
		return nil, fmt.Errorf("database type not supported: %s", conns.Databases.Type)
	}
	articleLogger := logger.WithField("package", "article")
	articleSrv := article.NewService(articleRepo, articleLogger)
	articleHandler := articlehandler.New(articleSrv, articleLogger)

	indexWebHandler := indexhandler.New()
	routes := []route{
		// infra
		{
			path:    "/readiness",
			method:  get,
			handler: health.Readiness().ServeHTTP,
		},
		// services
		{
			path:    "/api/v1/articles",
			method:  get,
			handler: articleHandler.List,
		},
		{
			path:    "/api/v1/articles/{id:[a-z0-9-\\-]+}",
			method:  post,
			handler: articleHandler.View,
		},
		// web
		{
			path:    "/",
			method:  get,
			handler: indexWebHandler.Index,
		},
	}

	loggingMW := middleware.Logging(logger.WithField("package", "middleware"))
	r := mux.NewRouter()
	r.Use(middleware.EnableCORS)
	r.Use(middleware.RequestID)
	r.Use(middleware.StatusResponseWriter)
	r.Use(loggingMW)
	r.Use(handlers.CompressHandler)

	for _, rt := range routes {
		h := rt.handler
		for i := len(rt.middlewares) - 1; i >= 0; i-- {
			h = rt.middlewares[i](h)
		}
		r.Path(rt.path).Methods(rt.method).HandlerFunc(h)
	}

	// even not found, return index so that VueJS does its job
	r.NotFoundHandler = middleware.RequestID(loggingMW(http.HandlerFunc(indexWebHandler.Index)))

	// static resources
	static := []struct {
		prefix string
		dir    string
	}{
		{
			prefix: "/",
			dir:    "web/",
		},
	}
	for _, s := range static {
		h := http.StripPrefix(s.prefix, http.FileServer(http.Dir(s.dir)))
		r.PathPrefix(s.prefix).Handler(middleware.StaticCache(h, 3600*24))
	}

	return r, nil
}

// Close close all underlying connections
func (c *InfraConns) Close() {
	c.Databases.Close()
}
