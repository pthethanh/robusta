package api

import (
	"github.com/pthethanh/robusta/internal/app/notification"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func createNotificationService(es event.Subscriber, user notification.UserService) (*notification.Service, *closeutil.Closer, error) {
	mailer := createMailer()
	conf := notification.LoadConfigFromEnv()
	srv, err := notification.NewService(conf, mailer, es, user)
	if err != nil {
		return nil, nil, err
	}
	closer := closeutil.NewCloser()
	return srv, closer, nil
}
