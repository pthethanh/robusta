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
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	Service struct {
		conf   Config
		mailer email.Sender
		es     event.Subscriber
		wg     sync.WaitGroup
	}

	Config struct {
		CommentTopic  string        `envconfig:"COMMENT_TOPIC" default:"r_topic_comment"`
		ReactionTopic string        `envconfig:"REACTION_TOPIC" default:"r_topic_reaction"`
		Worker        int           `envconfig:"NOTIFICATION_WORKER" default:"10"`
		CloseTimeout  time.Duration `envconfig:"NOTIFICATION_CLOSE_TIMEOUT" default:"15s"`
	}
)

func NewService(conf Config, mailer email.Sender, es event.Subscriber) *Service {
	s := &Service{
		conf:   conf,
		mailer: mailer,
		es:     es,
	}
	return s
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func (s *Service) Start() {
	ch := s.es.Subscribe(s.conf.CommentTopic, s.conf.ReactionTopic)
	for i := 0; i < s.conf.Worker; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for ev := range ch {
				switch ev.Type {
				case types.EventCommentCreated:
					// TODO implement me
					log.Infof("--------------comment created")
				case types.EventReactionCreated:
					// TODO implement me
					log.Infof("--------------reaction created")
				}
			}
		}()
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
}
