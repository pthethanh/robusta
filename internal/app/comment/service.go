package comment

import (
	"context"
	"fmt"
	"sync"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/uuid"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	Repository interface {
		Create(ctx context.Context, a *types.Comment) error
		FindAll(ctx context.Context, req FindRequest) ([]*types.Comment, error)
		UpdateReactions(ctx context.Context, id string, req *types.ReactionDetail) error
		Update(ctx context.Context, id string, a *types.Comment) error
		Delete(ctx context.Context, id string) (types.Comment, error)
		FindByID(ctx context.Context, id string) (types.Comment, error)
	}

	PolicyService interface {
		Validate(ctx context.Context, obj string, act string) error
		AddPolicy(ctx context.Context, p types.Policy) error
	}

	Config struct {
		Topic             string `envconfig:"COMMENT_TOPIC" default:"r_topic_comment"`
		ReactionTopic     string `envconfig:"REACTION_TOPIC" default:"r_topic_reaction"`
		NotificationTopic string `envconfig:"NOTIFICATION_TOPIC" default:"r_topic_notification"`
		EventWorkers      int    `envconfig:"COMMENT_EVENT_WORKERS" default:"10"`
		MaxPageSize       int    `envconfig:"COMMENT_MAX_PAGE_SIZE" default:"100"`
	}

	Service struct {
		conf   Config
		repo   Repository
		policy PolicyService
		es     event.PublishSubscriber
		wait   sync.WaitGroup
	}
)

func NewService(conf Config, repo Repository, policy PolicyService, es event.PublishSubscriber) *Service {
	srv := &Service{
		conf:   conf,
		repo:   repo,
		policy: policy,
		es:     es,
	}
	go func() {
		srv.listenEvents()
	}()
	return srv
}

// LoadConfigFromEnv config from environment variables
func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func (s *Service) Create(ctx context.Context, cm *types.Comment) error {
	if err := validator.Validate(cm); err != nil {
		return fmt.Errorf("invalid comment: %w", err)
	}
	user := auth.FromContext(ctx)
	if user != nil {
		cm.CreatedByID = user.UserID
		cm.CreatedByName = user.GetName()
		cm.CreatedByAvatar = user.AvatarURL
	}
	// thread_id only available for level 0
	// a reply to comment should inherit the thread it from its parent
	// this make the entire tree has same thread_id with the root node
	// this is for easier to group comments as a thread in the future.
	if cm.Level == 0 && cm.ThreadID == "" {
		cm.ThreadID = uuid.New()
	}
	if err := s.repo.Create(ctx, cm); err != nil {
		log.WithContext(ctx).Errorf("failed to create comment, err: %v", err)
		return err
	}
	// make the user the owner of the comment
	if err := s.policy.AddPolicy(auth.NewAdminContext(ctx), types.Policy{
		Subject: user.UserID,
		Object:  cm.ID,
		Action:  types.PolicyActionAny,
		Effect:  types.PolicyEffectAllow,
	}); err != nil {
		log.WithContext(ctx).Errorf("failed to make owner for the comment, err: %v", err)
		return err
	}
	// send event to who interested in
	defer s.sendEvent(ctx, *cm, types.EventCommentCreated)

	// send reply notification to the parent comment's owner
	if cm.ReplyToID != "" {
		s.sendReplyCreatedNotification(*cm)
	}
	return nil
}

func (s *Service) FindAll(ctx context.Context, req FindRequest) ([]*types.Comment, error) {
	if req.Limit > s.conf.MaxPageSize {
		req.Limit = s.conf.MaxPageSize
	}
	return s.repo.FindAll(ctx, req)
}

func (s *Service) Update(ctx context.Context, id string, cm *types.Comment) error {
	if err := validator.Validate(cm); err != nil {
		return fmt.Errorf("invalid comment: %w", err)
	}
	if err := s.policy.Validate(ctx, id, types.PolicyActionCommentUpdate); err != nil {
		return err
	}
	return s.repo.Update(ctx, id, cm)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.policy.Validate(ctx, id, types.PolicyActionCommentUpdate); err != nil {
		return err
	}
	deletedComment, err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	defer s.sendEvent(ctx, deletedComment, types.EventCommentDeleted)
	return nil
}

func (s *Service) FindByID(ctx context.Context, id string) (types.Comment, error) {
	return s.repo.FindByID(ctx, id)
}

// Close close/wait underlying services
func (s *Service) Close() error {
	// wait for background processes finish their jobs
	s.wait.Wait()
	return nil
}
