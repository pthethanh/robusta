package challenge

import (
	"context"
	"fmt"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	PolicyService interface {
		Validate(ctx context.Context, obj string, act string) error
		AddPolicy(ctx context.Context, p types.Policy) error
	}

	Repository interface {
		Insert(ctx context.Context, c *types.Challenge) error
		FindByID(ctx context.Context, id string) (*types.Challenge, error)
		FindAll(ctx context.Context, r FindRequest) ([]*types.Challenge, error)
		Delete(cxt context.Context, id string) error
		Update(cxt context.Context, req UpdateRequest) error
	}
	Service struct {
		conf   Config
		repo   Repository
		policy PolicyService
	}

	Config struct {
		MaxPageSize int `envconfig:"CHALLENGE_MAX_PAGE_SIZE" default:"50"`
	}
)

func NewService(conf Config, repo Repository, policy PolicyService) *Service {
	return &Service{
		conf:   conf,
		repo:   repo,
		policy: policy,
	}
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

// Create create a challenge.
// For now everyone is allowed to create challenge, but this can be changed in the future.
func (s *Service) Create(ctx context.Context, c *types.Challenge) error {
	if err := validator.Validate(c); err != nil {
		log.WithContext(ctx).Errorf("invalid input, err: %v", err)
		return status.Gen().BadRequest
	}
	if err := s.policy.Validate(ctx, types.PolicyObjectChallenge, types.PolicyActionChallengeCreate); err != nil {
		log.WithContext(ctx).Errorf("failed to create challenge due to permission issue, err: %v", err)
		return err
	}
	user := auth.FromContext(ctx)
	if user != nil {
		c.CreatedByID = user.UserID
		c.CreatedByName = user.GetName()
		c.CreatedByAvatar = user.AvatarURL
	}
	if err := s.repo.Insert(ctx, c); err != nil {
		return fmt.Errorf("failed to insert challenge: %w", err)
	}
	if err := s.policy.AddPolicy(auth.NewAdminContext(ctx), types.Policy{
		Subject: user.UserID,
		Object:  c.ID,
		Action:  types.PolicyActionAny,
		Effect:  types.PolicyEffectAllow,
	}); err != nil {
		return fmt.Errorf("failed to set permission: %w", err)
	}
	return nil
}

// Get return full information of the requested challenge.
// The requested user must have read permission, otherwise the request will be rejected.
func (s *Service) Get(ctx context.Context, id string) (*types.Challenge, error) {
	if err := s.policy.Validate(ctx, id, types.PolicyActionChallengeRead); err != nil {
		log.WithContext(ctx).Errorf("failed to read challenge due to permission issue, err: %v", err)
		return nil, err
	}
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find the challenge: %w", err)
	}
	return c, nil
}

// FindAll return list of challenges base on the given request without their test detail.
// To prevent leak of private challenges, a folder_id  (group) must be provided.
// And the request user must have permission on the folder to be able to get list of challenges of the folder (group).
func (s *Service) FindAll(ctx context.Context, r FindRequest) ([]*types.Challenge, error) {
	if err := validator.Validate(r); err != nil {
		return nil, err
	}
	if r.Limit > s.conf.MaxPageSize {
		r.Limit = s.conf.MaxPageSize
	}
	// if query across folders, user must have read-list permission.
	// otherwise he/she must have read permission on the given folder.
	if r.FolderID == "" {
		if err := s.policy.Validate(ctx, types.PolicyObjectFolder, types.PolicyActionFolderReadList); err != nil {
			log.WithContext(ctx).Errorf("reading list of challenges failed due to permission issue, err: %v", err)
			return nil, err
		}
	} else if err := s.policy.Validate(ctx, r.FolderID, types.PolicyActionFolderRead); err != nil {
		log.WithContext(ctx).Errorf("reading list of challenges failed due to permission issue, err: %v", err)
		return nil, err
	}

	challenges, err := s.repo.FindAll(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to find challenges: %w", err)
	}
	return challenges, nil
}

// Delete delete the given challenge.
func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.policy.Validate(ctx, id, types.PolicyActionChallengeDelete); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

// FindFolderChallengeByID find the challenge by the given id.
// User must have folder:read permission to be able to get it.
func (s *Service) FindFolderChallengeByID(ctx context.Context, id string, folderID string) (*types.Challenge, error) {
	if err := s.policy.Validate(ctx, folderID, types.PolicyActionFolderRead); err != nil {
		log.WithContext(ctx).Errorf("reading list of challenges failed due to permission issue, err: %v", err)
		return nil, err
	}
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find the challenge: %w", err)
	}
	return c, nil
}

func (s *Service) Update(ctx context.Context, req UpdateRequest) error {
	if err := validator.Validate(req); err != nil {
		log.WithContext(ctx).Errorf("bad request, err: %v", err)
		return status.Gen().BadRequest
	}
	if err := s.policy.Validate(ctx, req.ID, types.PolicyActionChallengeUpdate); err != nil {
		log.WithContext(ctx).Errorf("failed to update challenge due to permission issue, err: %v", err)
		return err
	}
	if err := s.repo.Update(ctx, req); err != nil {
		log.WithContext(ctx).Errorf("failed to update challenge, err: %v", err)
		return fmt.Errorf("failed to update challenge: %w", err)
	}
	return nil
}
