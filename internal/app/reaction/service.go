package reaction

import (
	"context"
	"fmt"
	"net/url"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	repository interface {
		Upsert(ctx context.Context, reaction *types.Reaction) (bool, *types.Reaction, error)
		Find(ctx context.Context, queries map[string]interface{}) (types.Reactions, error)
	}

	Config struct {
		Topic string `envconfig:"REACTION_TOPIC" default:"r_topic_reaction"`
	}

	Service struct {
		conf       Config
		repo       repository
		eventStore event.Publisher
	}
)

func NewService(conf Config, repo repository, es event.Publisher) *Service {
	return &Service{
		conf:       conf,
		repo:       repo,
		eventStore: es,
	}
}

// LoadConfigFromEnv config from environment variables
func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

// Create create a reaction if not exist, otherwise update the existing one
// A reaction is determined as already exist if there is a reaction with same:
// - target_id
// - target_type
// - created_by_id
func (s *Service) Create(ctx context.Context, reaction *types.Reaction) (*types.ReactionDetail, error) {
	if err := validator.Validate(reaction); err != nil {
		return nil, fmt.Errorf("invalid comment: %w", err)
	}
	user := auth.FromContext(ctx)
	if user != nil {
		reaction.CreatedByID = user.UserID
		reaction.CreatedByName = user.GetName()
		reaction.CreatedByAvatar = user.AvatarURL
	}
	created, oldReaction, err := s.repo.Upsert(ctx, reaction)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to create reaction, err: %v", err)
		return nil, err
	}
	log.WithContext(ctx).WithFields(log.Fields{
		"reaction_is_new": created,
		"target_id":       reaction.ID,
		"created_by_id":   user.UserID,
	}).Debug("new reaction is created")
	// collect current reactions
	queries := fmt.Sprintf("target_type=%s&target_id=%s", reaction.TargetType, reaction.TargetID)
	reactions, err := s.Find(ctx, queries)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to find reactions, err: %v", err)
		return nil, err
	}
	// send event to who interested in
	defer s.sendEvent(ctx, &types.ReactionChanged{
		IsNew:       created,
		OldReaction: oldReaction,
		NewReaction: reaction,
		Detail:      reactions.Detail(),
	}, types.EventReactionCreated)
	return reactions.Detail(), nil
}

// Find find the reactions base on the given queries.
// The given queries must be in form of HTTP query
func (s *Service) Find(ctx context.Context, queries string) (types.Reactions, error) {
	recorder := timeutil.NewRecorder("find by target")
	defer log.WithContext(ctx).Info(recorder)

	log.WithContext(ctx).Debugf("input queries: %v", queries)
	values, err := url.ParseQuery(queries)
	if err != nil {
		return nil, err
	}
	q := make(map[string]interface{})
	if _, ok := values["target_id"]; ok {
		q["target_id"] = values.Get("target_id")
	}
	if _, ok := values["target_type"]; ok {
		q["target_type"] = values.Get("target_type")
	}
	if _, ok := values["created_by_id"]; ok {
		q["created_by_id"] = values.Get("created_by_id")
	}
	if _, ok := values["id"]; ok {
		q["_id"] = values.Get("id")
	}
	if _, ok := values["type"]; ok {
		q["type"] = values.Get("type")
	}

	return s.repo.Find(ctx, q)
}
