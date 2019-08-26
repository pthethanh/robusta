package notification

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/email"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	Service struct {
		conf   Config
		mailer email.Sender
		ch     chan types.Notification
		wg     sync.WaitGroup
	}

	Config struct {
		Buffer       int           `envconfig:"NOTIFICATION_BUFFER" default:"1000"`
		Worker       int           `envconfig:"NOTIFICATION_WORKER" default:"10"`
		CloseTimeout time.Duration `envconfig:"NOTIFICATION_CLOSE_TIMEOUT" default:"15s"`
	}
)

func NewService(conf Config, mailer email.Sender) *Service {
	s := &Service{
		conf:   conf,
		mailer: mailer,
		ch:     make(chan types.Notification),
	}
	// start worker pool
	s.start()
	return s
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func (s *Service) Notify(ctx context.Context, info types.Notification) {
	s.ch <- info
}

func (s *Service) start() {
	for i := 0; i < s.conf.Worker; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for info := range s.ch {
				if err := s.sendNotification(info); err != nil {
					log.Errorf("failed to send notification, err: %v", err)
					continue
				}
				log.Infof("sent notification, type: %s", info.Type)
			}
		}()
	}
}

func (s *Service) sendNotification(info types.Notification) error {
	switch info.Type {
	case types.NotificationTypeEmail:
		return s.sendEmail(info)
	default:
		return ErrNotSupported
	}
}

func (s *Service) sendEmail(info types.Notification) error {
	var m email.Email
	if err := json.Unmarshal(info.Data, &m); err != nil {
		return err
	}
	if err := s.mailer.Send(context.Background(), m); err != nil {
		return errors.Wrap(err, "failed to send email")
	}
	return nil
}

func (s *Service) Close() error {
	close(s.ch)
	done := make(chan struct{})
	go func() {
		log.Infof("notification: waiting for workers to be shutdown....")
		s.wg.Wait()
		done <- struct{}{}
	}()
	defer log.Infof("notification: bye bye...")
	select {
	case <-done:
		return nil
	case <-time.After(s.conf.CloseTimeout):
		return ErrTimeout
	}
	return nil
}
