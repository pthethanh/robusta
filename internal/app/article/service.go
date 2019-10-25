package article

import (
	"context"
	"fmt"
	"sync"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	// Repository is an interface of an article repository
	Repository interface {
		FindAll(ctx context.Context, req FindRequest) ([]*Article, error)
		Increase(ctx context.Context, id string, field string, val interface{}) error
		Create(ctx context.Context, a *Article) error
		FindByID(ctx context.Context, id string) (*Article, error)
		FindByArticleID(ctx context.Context, id string) (*Article, error)
		ChangeStatus(ctx context.Context, id string, status Status) error
		Update(ctx context.Context, id string, a *Article) error
		UpdateReactions(ctx context.Context, id string, req *types.ReactionDetail) error
	}

	PolicyService interface {
		Validate(ctx context.Context, obj string, act string) error
		AddPolicy(ctx context.Context, p types.Policy) error
	}

	Config struct {
		CommentTopic      string `envconfig:"COMMENT_TOPIC" default:"r_topic_comment"`
		ReactionTopic     string `envconfig:"REACTION_TOPIC" default:"r_topic_reaction"`
		NotificationTopic string `envconfig:"NOTIFICATION_TOPIC" default:"r_topic_notification"`
		EventWorkers      int    `envconfig:"ARTICLE_EVENT_WORKERS" default:"10"`
		MaxPageSize       int    `envconfig:"ARTICLE_MAX_PAGE_SIZE" default:"15"`
	}

	// Service is an article Service
	Service struct {
		conf   Config
		repo   Repository
		policy PolicyService
		es     event.PublishSubscriber
		wait   sync.WaitGroup
	}
)

// NewService return a new article service
func NewService(conf Config, r Repository, policySrv PolicyService, es event.PublishSubscriber) *Service {
	srv := &Service{
		conf:   conf,
		repo:   r,
		policy: policySrv,
		es:     es,
	}

	// handling events from other services
	go func() {
		srv.handleEvents()
	}()
	return srv
}

// LoadConfigFromEnv config from environment variables
func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

// FindAll return all articles
func (s *Service) FindAll(ctx context.Context, req FindRequest) ([]*Article, error) {
	recorder := timeutil.NewRecorder("find all articles")
	defer log.WithContext(ctx).Info(recorder)
	if err := validator.Validate(req); err != nil {
		return nil, fmt.Errorf("invalid find request: %w", err)
	}
	if req.Limit > s.conf.MaxPageSize {
		req.Limit = s.conf.MaxPageSize
	}
	articles, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to find articles: %w", err)
	}
	return articles, nil
}

// Create create a new article
func (s *Service) Create(ctx context.Context, a *Article) error {
	if err := validator.Validate(a); err != nil {
		return fmt.Errorf("invalid article: %w", err)
	}

	a.Status = StatusPublic
	user := auth.FromContext(ctx)
	if user != nil {
		a.CreatedByID = user.UserID
		a.CreatedByName = user.GetName()
		a.CreatedByAvatar = user.AvatarURL
	}
	a.ArticleID = articleID(a.Title)
	if err := s.repo.Create(ctx, a); err != nil {
		log.WithContext(ctx).Errorf("failed to create article, err: %v", err)
		return fmt.Errorf("failed to insert article:  %w", err)
	}

	// make her the owner of the article
	if err := s.policy.AddPolicy(auth.NewAdminContext(ctx), types.Policy{
		Subject: user.UserID,
		Object:  a.ID,
		Action:  types.PolicyActionAny,
		Effect:  types.PolicyEffectAllow,
	}); err != nil {
		return err
	}
	return nil
}

// ChangeStatus delete the given article
func (s *Service) ChangeStatus(ctx context.Context, id string, status Status) error {
	if err := s.policy.Validate(ctx, id, types.PolicyActionArticleUpdate); err != nil {
		return err
	}
	return s.repo.ChangeStatus(ctx, id, status)
}

// Update the existing article
func (s *Service) Update(ctx context.Context, id string, a *Article) error {
	if err := validator.Validate(a); err != nil {
		return fmt.Errorf("invalid article: %w", err)
	}
	if err := s.policy.Validate(ctx, id, types.PolicyActionArticleUpdate); err != nil {
		return err
	}
	return s.repo.Update(ctx, id, a)
}

// FindByID find article by id
func (s *Service) FindByID(ctx context.Context, id string) (*Article, error) {
	a, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// IncreaseView increase number of view of the given article
func (s *Service) IncreaseView(ctx context.Context, id string) error {
	return s.repo.Increase(ctx, id, "views", 1)
}

// FindByArticleID find article by id
func (s *Service) FindByArticleID(ctx context.Context, articleID string) (*Article, error) {
	a, err := s.repo.FindByArticleID(ctx, articleID)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Close close/wait underlying background process to finish
func (s *Service) Close() error {
	// wait for background process to finished their job
	// see listenEvents for detail
	s.wait.Wait()
	return nil
}
