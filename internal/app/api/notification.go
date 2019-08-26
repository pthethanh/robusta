package api

import (
	"github.com/pthethanh/robusta/internal/app/notification"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func createNotificationService() (*notification.Service, *closeutil.Closer) {
	mailer := createMailer()
	conf := notification.LoadConfigFromEnv()
	srv := notification.NewService(conf, mailer)
	closer := closeutil.NewCloser()
	closer.Add(srv.Close)
	return srv, closer
}
