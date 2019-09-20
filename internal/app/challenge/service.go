package challenge

import (
	"context"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/policy"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	PolicyService interface {
		IsAllowed(ctx context.Context, sub string, obj string, act string) bool
		MakeOwner(ctx context.Context, sub string, obj string) error
	}

	Repository interface {
		Insert(ctx context.Context, c *types.Challenge) error
		FindByID(ctx context.Context, id string) (*types.Challenge, error)
		FindAll(ctx context.Context, r FindRequest) ([]*types.Challenge, error)
		Delete(cxt context.Context, id string) error
	}
	Service struct {
		repo   Repository
		policy PolicyService
	}
)

func NewService(repo Repository, policy PolicyService) *Service {
	return &Service{
		repo:   repo,
		policy: policy,
	}
}

// Create create a challenge.
// For now everyone is allowed to create challenge, but this can be changed in the future.
func (s *Service) Create(ctx context.Context, c *types.Challenge) error {
	if err := validator.Validate(c); err != nil {
		log.WithContext(ctx).Errorf("invalid input, err: %v", err)
		return types.ErrBadRequest
	}
	user := auth.FromContext(ctx)
	if user != nil {
		c.CreatedByID = user.UserID
		c.CreatedByName = user.GetName()
		c.CreatedByAvatar = user.AvatarURL
	}
	if err := s.repo.Insert(ctx, c); err != nil {
		return errors.Wrap(err, "failed to insert challenge")
	}
	if err := s.policy.MakeOwner(ctx, user.UserID, c.ID); err != nil {
		return errors.Wrap(err, "failed to set permission")
	}
	return nil
}

// Get return full information of the requested challenge.
// The requested user must have read permission, otherwise the request will be rejected.
func (s *Service) Get(ctx context.Context, id string) (*types.Challenge, error) {
	if err := s.isAllowed(ctx, id, ActionRead); err != nil {
		log.WithContext(ctx).Errorf("failed to read challenge due to permission issue, err: %v", err)
		return nil, err
	}
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find the challenge")
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
	if err := s.isAllowed(ctx, r.FolderID, ActionRead); err != nil {
		log.WithContext(ctx).Errorf("reading list of challenges failed due to permission issue, err: %v", err)
		return nil, err
	}
	challenges, err := s.repo.FindAll(ctx, r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find challenges")
	}
	return challenges, nil
}

// Delete delete the given challenge.
func (s *Service) Delete(ctx context.Context, id string) error {
	if err := policy.IsCurrentUserAllowed(ctx, s.policy, id, ActionDelete); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *Service) isAllowed(ctx context.Context, id string, action string) error {
	return policy.IsCurrentUserAllowed(ctx, s.policy, id, action)
}
