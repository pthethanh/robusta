package friend

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/glog"
)

// Repository is an interface of a friend repository
type Repository interface {
	FindByID(ctx context.Context, id string) (*types.Friend, error)
}

// Service is an friend service
type Service struct {
	repo   Repository
	logger glog.Logger
}

// NewService return a new friend service
func NewService(r Repository, l glog.Logger) *Service {
	return &Service{
		repo:   r,
		logger: l,
	}
}

// Get return given friend by his/her id
func (s *Service) Get(ctx context.Context, id string) (*types.Friend, error) {
	return s.repo.FindByID(ctx, id)
}
