package notification

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"sync"
	"time"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/email"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/htmlutil"
)

type (
	UserService interface {
		FindByUserID(ctx context.Context, id string) (*types.User, error)
	}
	Service struct {
		conf      Config
		mailer    email.Sender
		es        event.Subscriber
		wg        sync.WaitGroup
		templates *template.Template
		user      UserService
	}

	Config struct {
		Topic        string        `envconfig:"NOTIFICATION_TOPIC" default:"r_topic_notification"`
		Worker       int           `envconfig:"NOTIFICATION_WORKER" default:"10"`
		CloseTimeout time.Duration `envconfig:"NOTIFICATION_CLOSE_TIMEOUT" default:"15s"`
		TemplatePath string        `envconfig:"NOTIFICATION_TEMPLATE_PATH" default:"templates/notifications"`
	}
)

func NewService(conf Config, mailer email.Sender, es event.Subscriber, user UserService) (*Service, error) {
	templates, err := htmlutil.LoadTemplates(conf.TemplatePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load notification templates: %w", err)
	}
	s := &Service{
		conf:      conf,
		mailer:    mailer,
		es:        es,
		templates: templates,
		user:      user,
	}
	return s, nil
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func (s *Service) Start() {
	ch := s.es.Subscribe(s.conf.Topic)
	for i := 0; i < s.conf.Worker; i++ {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for ev := range ch {
				switch ev.Type {
				case types.EventNotificationArticleCommentCreated:
					s.handleArticleCommentCreated(ev)
				case types.EventNotificationArticleReactionCreated:
					s.handleArticleReactionCreated(ev)
				case types.EventNotificationCommentReactionCreated:
					s.handleCommentReactionCreated(ev)
				case types.EventNotificationCommentReplyCreated:
					s.handleCommentReplyCreated(ev)
				case types.EventPasswordResetTokenCreated:
					s.handleUserPasswordResetTokenCreated(ev)
				}
			}
		}()
	}
}

func (s *Service) sendEmailNotification(subject string, bodyTemplate string, bodyData interface{}, userID string) {
	owner, err := s.user.FindByUserID(context.Background(), userID)
	if err != nil {
		log.Errorf("failed to find user: %s, err: %v", userID, err)
		return
	}
	tmpl := s.templates.Lookup(bodyTemplate)
	if tmpl == nil {
		log.Errorf("failed to find the target template, err: %v", err)
		return
	}
	buff := bytes.Buffer{}
	if err := tmpl.Execute(&buff, bodyData); err != nil {
		log.Errorf("failed to execute template, err: %v", err)
		return
	}
	if err := s.mailer.Send(context.Background(), email.Email{
		Subject: subject,
		To:      []string{owner.Email},
		Body:    buff.String(),
	}); err != nil {
		log.Errorf("failed to send email, err :%v", err)
		return
	}
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
		return status.Gen().Timeout
	}
}
