package tutorial

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/policy"
)

type (
	PolicyService interface {
		IsAllowed(ctx context.Context, sub string, obj string, act string) bool
		MakeOwner(ctx context.Context, sub string, obj string) error
	}

	Repository interface {
		FindAll(ctx context.Context, offset, limit int) ([]*Tutorial, error)
		Increase(ctx context.Context, id string, field string, val interface{}) error
		Create(ctx context.Context, a *Tutorial) error
		FindByID(ctx context.Context, id string) (*Tutorial, error)
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, a *Tutorial) error
	}

	// Service is actually internal service but with permission checking
	// that should be used by public services that expose API to external usages...
	Service struct {
		repo   Repository
		policy PolicyService
	}
)

// NewService return new instance of service
func NewService(repo Repository, policy PolicyService) *Service {
	return &Service{
		repo:   repo,
		policy: policy,
	}
}

// FindAll return all tutorials
func (s *Service) FindAll(ctx context.Context, offset, limit int) ([]*Tutorial, error) {
	tutorials, err := s.repo.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all tutorials")
	}
	for i, a := range tutorials {
		if a.CreatedByName == "" && a.CreatedByID == "" {
			tutorials[i].CreatedByName = "goway"
		}
	}
	return tutorials, nil
}

// Create create a new tutorial
func (s *Service) Create(ctx context.Context, a *Tutorial) error {
	a.Status = StatusPublished // TODO will change in the future...
	if user := auth.FromContext(ctx); user != nil {
		a.CreatedByID = user.UserID
		a.CreatedByName = user.GetName()
		a.CreatedByAvatar = user.AvatarURL
	}
	if err := s.repo.Create(ctx, a); err != nil {
		return err
	}
	authUser := auth.FromContext(ctx)
	if err := s.policy.MakeOwner(ctx, authUser.UserID, a.ID); err != nil {
		return err
	}
	return nil
}

// Delete delete the given tutorial
func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.isAllowed(ctx, id, ActionDelete); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}

// Update the existing tutorial
func (s *Service) Update(ctx context.Context, id string, a *Tutorial) error {
	if err := s.isAllowed(ctx, id, ActionUpdate); err != nil {
		return err
	}
	return s.repo.Update(ctx, id, a)
}

// FindByID find tutorial by id
func (s *Service) FindByID(ctx context.Context, id string) (*Tutorial, error) {
	a, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// don't allow to find by ID for deleted tutorial
	if a.Status != StatusPublished {
		return nil, ErrNotFound
	}
	return a, nil
}

// View increase number of view of the given tutorial
func (s *Service) View(ctx context.Context, id string) error {
	return s.repo.Increase(ctx, id, "views", 1)
}

func (s *Service) isAllowed(ctx context.Context, id string, act string) error {
	return policy.IsCurrentUserAllowed(ctx, s.policy, id, act)
}
