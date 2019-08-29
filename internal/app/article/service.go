package article

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/policy"
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
		FindAll(ctx context.Context, req FindRequest) ([]*types.Article, error)
		Increase(ctx context.Context, id string, field string, val interface{}) error
		Create(ctx context.Context, a *types.Article) error
		FindByID(ctx context.Context, id string) (*types.Article, error)
		FindByArticleID(ctx context.Context, id string) (*types.Article, error)
		ChangeStatus(ctx context.Context, id string, status types.Status) error
		Update(ctx context.Context, id string, a *types.Article) error
		UpdateReactions(ctx context.Context, id string, req *types.ReactionDetail) error
	}

	PolicyService interface {
		IsAllowed(ctx context.Context, sub policy.Subject, obj policy.Object, act policy.Action) bool
		MakeOwner(ctx context.Context, sub policy.Subject, obj policy.Object) error
	}

	Config struct {
		CommentTopic  string `envconfig:"COMMENT_TOPIC" default:"r_topic_comment"`
		ReactionTopic string `envconfig:"REACTION_TOPIC" default:"r_topic_reaction"`
		EventWorkers  int    `envconfig:"ARTICLE_EVENT_WORKERS" default:"10"`
	}

	// Service is an article Service
	Service struct {
		conf   Config
		repo   Repository
		policy PolicyService
		es     event.Subscriber
		wait   sync.WaitGroup
	}
)

// NewService return a new article service
func NewService(conf Config, r Repository, policySrv PolicyService, es event.Subscriber) *Service {
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
func (s *Service) FindAll(ctx context.Context, req FindRequest) ([]*types.Article, error) {
	recorder := timeutil.NewRecorder("find all articles")
	defer log.WithContext(ctx).Info(recorder)

	articles, err := s.repo.FindAll(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all articles")
	}
	return articles, nil
}

// Create create a new article
func (s *Service) Create(ctx context.Context, a *types.Article) error {
	if err := validator.Validate(a); err != nil {
		return errors.Wrap(err, "invalid article")
	}

	a.Status = types.StatusPublished
	user := auth.FromContext(ctx)
	if user != nil {
		a.CreatedByID = user.UserID
		a.CreatedByName = user.GetName()
		a.CreatedByAvatar = user.AvatarURL
	}
	a.ArticleID = articleID(a.Title)
	if err := s.repo.Create(ctx, a); err != nil {
		log.WithContext(ctx).Errorf("failed to create article, err: %v", err)
		return errors.Wrap(err, "failed to insert article")
	}

	// make her the owner of the article
	if err := s.policy.MakeOwner(ctx, policy.UserSubject(user.UserID), policy.ArticleObject(a.ID)); err != nil {
		return err
	}
	return nil
}

// ChangeStatus delete the given article
func (s *Service) ChangeStatus(ctx context.Context, id string, status types.Status) error {
	if err := s.isAllowed(ctx, id, policy.ActionUpdate); err != nil {
		return err
	}
	return s.repo.ChangeStatus(ctx, id, status)
}

// Update the existing article
func (s *Service) Update(ctx context.Context, id string, a *types.Article) error {
	if err := validator.Validate(a); err != nil {
		return errors.Wrap(err, "invalid article")
	}
	if err := s.isAllowed(ctx, id, policy.ActionUpdate); err != nil {
		return err
	}
	return s.repo.Update(ctx, id, a)
}

// FindByID find article by id
func (s *Service) FindByID(ctx context.Context, id string) (*types.Article, error) {
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
func (s *Service) FindByArticleID(ctx context.Context, articleID string) (*types.Article, error) {
	a, err := s.repo.FindByArticleID(ctx, articleID)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *Service) isAllowed(ctx context.Context, id string, act policy.Action) error {
	return policy.IsCurrentUserAllowed(ctx, s.policy, policy.ArticleObject(id), act)
}

// Close close/wait underlying background process to finish
func (s *Service) Close() error {
	// wait for background process to finished their job
	// see listenEvents for detail
	s.wait.Wait()
	return nil
}
