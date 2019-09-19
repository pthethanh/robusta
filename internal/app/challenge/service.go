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

func (s *Service) Get(ctx context.Context, id string) (*types.Challenge, error) {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find the challenge")
	}
	return c, nil
}

func (s *Service) FindAll(ctx context.Context, r FindRequest) ([]*types.Challenge, error) {
	if err := validator.Validate(r); err != nil {
		return nil, err
	}
	challenges, err := s.repo.FindAll(ctx, r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find challenges")
	}
	return challenges, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := policy.IsCurrentUserAllowed(ctx, s.policy, id, policy.ActionDelete); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
