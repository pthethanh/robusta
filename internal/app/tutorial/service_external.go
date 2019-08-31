package tutorial

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/policy"
)

type (
	PolicyService interface {
		IsAllowed(ctx context.Context, sub policy.Subject, obj policy.Object, act policy.Action) bool
		MakeOwner(ctx context.Context, sub policy.Subject, obj policy.Object) error
	}

	// ExternalService is actually internal service but with permission checking
	// that should be used by public services that expose API to external usages...
	ExternalService struct {
		internal *InternalService
		policy   PolicyService
	}
)

// NewExternalService return new instance of external service
func NewExternalService(internal *InternalService, policy PolicyService) *ExternalService {
	return &ExternalService{
		internal: internal,
		policy:   policy,
	}
}

// FindAll return all tutorials
func (s *ExternalService) FindAll(ctx context.Context, offset, limit int) ([]*Tutorial, error) {
	return s.internal.FindAll(ctx, offset, limit)
}

// Create create a new tutorial
func (s *ExternalService) Create(ctx context.Context, a *Tutorial) error {
	if err := s.internal.Create(ctx, a); err != nil {
		return err
	}
	authUser := auth.FromContext(ctx)
	if err := s.policy.MakeOwner(ctx, policy.UserSubject(authUser.UserID), policy.TutorialObject(a.ID)); err != nil {
		return err
	}
	return nil
}

// Delete delete the given tutorial
func (s *ExternalService) Delete(ctx context.Context, id string) error {
	if err := s.isAllowed(ctx, id, policy.ActionDelete); err != nil {
		return err
	}
	return s.internal.Delete(ctx, id)
}

// Update the existing tutorial
func (s *ExternalService) Update(ctx context.Context, id string, a *Tutorial) error {
	if err := s.isAllowed(ctx, id, policy.ActionUpdate); err != nil {
		return err
	}
	return s.internal.Update(ctx, id, a)
}

// FindByID find tutorial by id
func (s *ExternalService) FindByID(ctx context.Context, id string) (*Tutorial, error) {
	a, err := s.internal.FindByID(ctx, id)
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
func (s *ExternalService) View(ctx context.Context, id string) error {
	return s.internal.View(ctx, id)
}

func (s *ExternalService) isAllowed(ctx context.Context, id string, act policy.Action) error {
	return policy.IsCurrentUserAllowed(ctx, s.policy, policy.TutorialObject(id), act)
}
