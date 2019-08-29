package api

import (
	"github.com/pthethanh/robusta/internal/app/notification"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func createNotificationService(es event.Subscriber) (*notification.Service, *closeutil.Closer) {
	mailer := createMailer()
	conf := notification.LoadConfigFromEnv()
	srv := notification.NewService(conf, mailer, es)
	closer := closeutil.NewCloser()
	return srv, closer
}
