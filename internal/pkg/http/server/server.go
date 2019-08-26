package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/pthethanh/robusta/internal/pkg/log"
)

// ListenAndServe will start a HTTP server base on the given configurations.
// A HTTPS server will be started if TLS options are configured.
func ListenAndServe(conf Config, router http.Handler) {
	port := fmt.Sprint(conf.Port)
	if conf.Port == 0 {
		port = os.Getenv("PORT") // mostly for app engine or heroku
		if port == "" {
			port = "80"
		}
	}
	address := fmt.Sprintf("%s:%s", conf.Address, port)
	srv := &http.Server{
		Addr:              address,
		Handler:           router,
		ReadTimeout:       conf.ReadTimeout,
		WriteTimeout:      conf.WriteTimeout,
		ReadHeaderTimeout: conf.ReadHeaderTimeout,
	}
	log.Infof("HTTP Server is listening on: %s", address)
	go func() {
		if conf.TLSCertFile != "" && conf.TLSKeyFile != "" {
			if err := srv.ListenAndServeTLS(conf.TLSCertFile, conf.TLSKeyFile); err != nil && err != http.ErrServerClosed {
				log.Panicf("listen: %s\n", err)
			}
			return
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals

	srvCtx, srvCancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
	defer srvCancel()
	log.Infof("shutting down http server...")
	if err := srv.Shutdown(srvCtx); err != nil {
		log.Panicf("http server shutdown with error:", err)
	}
}
