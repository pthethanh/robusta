package api

import (
	"io"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/health"
	"github.com/pthethanh/robusta/internal/pkg/http/middleware"
	"github.com/pthethanh/robusta/internal/pkg/http/router"
	"github.com/pthethanh/robusta/internal/pkg/limiter"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

const (
	get     = http.MethodGet
	post    = http.MethodPost
	put     = http.MethodPut
	delete  = http.MethodDelete
	options = http.MethodOptions
)

// NewRouter return new router
func NewRouter() (http.Handler, io.Closer, error) {
	closer := closeutil.NewCloser()
	es := event.NewMemoryEventStore(event.LoadConfigFromEnv())
	closer.Add(es.Close)

	policySrv, err := newPolicyService()
	if err != nil {
		return nil, closer, err
	}
	policyHandler := newPolicyHandler(policySrv)

	userSrv, userCloser, err := newUserService(policySrv, es)
	if err != nil {
		return nil, closer, err
	}
	closer.Append(userCloser)
	userHandler := newUserHandler(userSrv)

	notifier, notifierCloser, err := createNotificationService(es, userSrv)
	if err != nil {
		return nil, closer, err
	}
	go notifier.Start()

	reactionHandler, reactionCloser, err := createReactionHandler(es)
	if err != nil {
		return nil, closer, err
	}
	closer.Append(reactionCloser)

	commentHandler, commentCloser, err := newCommentHandler(policySrv, es)
	if err != nil {
		return nil, closer, err
	}
	closer.Append(commentCloser)

	articleHandler, articleCloser, err := newArticleHandler(policySrv, es)
	if err != nil {
		return nil, closer, err
	}
	closer.Append(articleCloser)

	editorHandler, err := newEditorHandler()
	if err != nil {
		return nil, closer, err
	}

	challengeHandler, challengeSrv, challengeCloser, err := newChallengeHandler(policySrv)
	if err != nil {
		return nil, closer, err
	}
	closer.Append(challengeCloser)
	solutionHandler, solutionSrv, solutionCloser, err := newSolutionHandler(policySrv)
	if err != nil {
		return nil, closer, err
	}
	closer.Append(solutionCloser)
	playgroundHandler := newPlaygroundHandler(challengeSrv, solutionSrv)

	folderHandler, folderCloser, err := newFolderHandler(policySrv)
	if err != nil {
		return nil, closer, err
	}
	closer.Append(folderCloser)

	jwtSignVerifier := newJWTSignVerifier()
	oauthHandler := newOAuth2Handler(jwtSignVerifier, userSrv)
	userInfoMiddleware := auth.UserInfoMiddleware(jwtSignVerifier)

	authHandler := newAuthHandler(jwtSignVerifier, map[string]auth.Authenticator{
		"local": userSrv,
	})

	// close notifier after close other services as the other services might generate notification
	// while they are shutting down
	closer.Append(notifierCloser)

	rateLimiter := limiter.New(limiter.LoadConfigFromEnv())
	indexHandler := NewIndexHandler()
	routes := []router.Route{
		// infra
		{
			Path:    "/readiness",
			Method:  get,
			Handler: health.Readiness().ServeHTTP,
		},
		// web
		{
			Path:    "/",
			Method:  get,
			Handler: indexHandler.ServeHTTP,
		},
	}
	// services routes
	routes = append(routes, authHandler.Routes()...)
	routes = append(routes, oauthHandler.Routes()...)
	routes = append(routes, userHandler.Routes()...)
	routes = append(routes, articleHandler.Routes()...)
	routes = append(routes, editorHandler.Routes()...)
	routes = append(routes, playgroundHandler.Routes()...)
	routes = append(routes, commentHandler.Routes()...)
	routes = append(routes, reactionHandler.Routes()...)
	routes = append(routes, challengeHandler.Routes()...)
	routes = append(routes, solutionHandler.Routes()...)
	routes = append(routes, folderHandler.Routes()...)
	routes = append(routes, policyHandler.Routes()...)

	// setting up router
	conf := router.LoadConfigFromEnv()
	conf.Routes = routes
	conf.Middlewares = []router.Middleware{
		rateLimiter.Limit,
		middleware.Recover,
		userInfoMiddleware,
		middleware.StatusResponseWriter,
		log.NewHTTPContextHandler(log.Root()),
		middleware.HTTPRequestResponseInfo(nil),
		middleware.Compress, // TODO remember disable compress when using http push
	}
	// even not found, return index so that VueJS does its job
	conf.NotFoundHandler = indexHandler

	r, err := router.New(conf)
	if err != nil {
		return nil, closer, err
	}
	health.Ready()
	return middleware.CORS(r), closer, nil
}
