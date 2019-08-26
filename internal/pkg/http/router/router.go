package router

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/http/middleware"

	"github.com/gorilla/mux"
)

type (
	// Config hold configurations of router
	Config struct {
		StaticPaths       map[string]string `envconfig:"STATIC_PATHS" default:"/static/:web/dist"`
		StaticCacheMaxAge int               `envconfig:"STATIC_CACHE_MAX_AGE" default:"31536000"`

		// need to set manually
		Middlewares     []Middleware
		Routes          []Route
		NotFoundHandler http.Handler
	}

	// Middleware is HTTP middleware
	Middleware = func(http.Handler) http.Handler

	// Route hold configuration of routing
	Route struct {
		Desc        string
		Path        string
		Method      string
		Queries     []string
		Handler     http.HandlerFunc
		Middlewares []Middleware
	}
)

// New create new router from Config.
func New(conf *Config) (http.Handler, error) {
	r := mux.NewRouter()
	for _, middleware := range conf.Middlewares {
		r.Use(middleware)
	}

	for _, rt := range conf.Routes {
		var h http.Handler
		h = http.HandlerFunc(rt.Handler)
		for i := len(rt.Middlewares) - 1; i >= 0; i-- {
			h = rt.Middlewares[i](h)
		}
		r.Path(rt.Path).Methods(rt.Method).Handler(h).Queries(rt.Queries...)
	}

	for prefix, dir := range conf.StaticPaths {
		h := http.StripPrefix(prefix, http.FileServer(http.Dir(dir)))
		r.PathPrefix(prefix).Handler(middleware.StaticCache(h, conf.StaticCacheMaxAge))
	}

	if conf.NotFoundHandler != nil {
		r.NotFoundHandler = conf.NotFoundHandler
	}
	return r, nil
}

// LoadConfigFromEnv load router configurations from environment variables
func LoadConfigFromEnv() *Config {
	var conf Config
	envconfig.Load(&conf)
	return &conf
}
