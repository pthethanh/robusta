package tutorial

import (
	"context"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/auth"
)

type (
	// Repository is an interface of an tutorial repository
	Repository interface {
		FindAll(ctx context.Context, offset, limit int) ([]*Tutorial, error)
		Increase(ctx context.Context, id string, field string, val interface{}) error
		Create(ctx context.Context, a *Tutorial) error
		FindByID(ctx context.Context, id string) (*Tutorial, error)
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, a *Tutorial) error
	}

	// InternalService is an tutorial InternalService
	InternalService struct {
		repo Repository
	}
)

// NewInternalService return a new tutorial InternalService
func NewInternalService(r Repository) *InternalService {
	return &InternalService{
		repo: r,
	}
}

// FindAll return all tutorials
func (s *InternalService) FindAll(ctx context.Context, offset, limit int) ([]*Tutorial, error) {
	tutorials, err := s.repo.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all tutorials")
	}
	for i, a := range tutorials {
		if a.CreatedByName == "" && a.CreatedByID == "" {
			tutorials[i].CreatedByName = "goway"
		}
		if a.Source == "" {
			tutorials[i].Source = "goway"
		}
	}
	return tutorials, nil
}

// Create create a new tutorial
func (s *InternalService) Create(ctx context.Context, a *Tutorial) error {
	a.Status = StatusPublished // TODO will change in the future...
	if user := auth.FromContext(ctx); user != nil {
		a.CreatedByID = user.UserID
		a.CreatedByName = user.GetName()
	}
	return s.repo.Create(ctx, a)
}

// Delete delete the given tutorial
func (s *InternalService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// Update the existing tutorial
func (s *InternalService) Update(ctx context.Context, id string, a *Tutorial) error {
	return s.repo.Update(ctx, id, a)
}

// FindByID find tutorial by id
func (s *InternalService) FindByID(ctx context.Context, id string) (*Tutorial, error) {
	a, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// View increase number of view of the given tutorial
func (s *InternalService) View(ctx context.Context, id string) error {
	return s.repo.Increase(ctx, id, "views", 1)
}
