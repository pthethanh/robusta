package article

import (
	"context"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/glog"
)

// Repository is an interface of an article repository
type Repository interface {
	FindAll(ctx context.Context, offset, limit int) ([]*types.Article, error)
}

// Service is an article service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new article service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// FindAll return all articles
func (s *Service) FindAll(ctx context.Context, offset, limit int) ([]*types.Article, error) {
	articles, err := s.repo.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all articles")
	}
	return articles, nil
}

// Create create a new article
func (s *Service) Create(ctx context.Context, a *types.Article) error {
	return nil
}

// Delete delete the given article
func (s *Service) Delete(ctx context.Context, id string) error {
	return nil
}

// Update the existing article
func (s *Service) Update(ctx context.Context, id string, a *types.Article) error {
	return nil
}

func (s *Service) Get(ctx context.Context, id string) (*types.Article, error) {
	return nil, nil
}
