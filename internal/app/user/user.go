package user

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/glog"

	"github.com/pkg/errors"
)

type (
	Repository interface {
		Create(context.Context, *types.User) (string, error)
		Delete(context.Context, string) error
		Update(context.Context, *types.User) error
		Lock(context.Context, string) error
		FindBySample(context.Context, *types.User) ([]*types.User, error)
		FindAll(context.Context) ([]*types.User, error)
	}
	PolicyService interface {
		IsAllowed(ctx context.Context, subject []string, action, resource string) (bool, error)
	}
	Service struct {
		logger glog.Logger
		policy PolicyService
		repo   Repository
	}
)

var (
	ActionCreate        = "user.create"
	ActionUpdate        = "user.update"
	ActionDelete        = "user.delete"
	ActionLock          = "user.lock"
	ActionReadList      = "user.read.list"
	PermissionResource  = "user"
	ErrPermissionDenied = errors.New("permission denied")
)

func (s *Service) Create(ctx context.Context, user *types.User) (string, error) {
	if err := s.IsAllowed(ctx, ActionCreate); err != nil {
		return "", err
	}
	id, err := s.repo.Create(ctx, user)
	return id, err
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.IsAllowed(ctx, ActionDelete); err != nil {
		return err
	}
	err := s.repo.Delete(ctx, id)
	return err
}

func (s *Service) Update(ctx context.Context, user *types.User) error {
	if err := s.IsAllowed(ctx, ActionUpdate); err != nil {
		return err
	}
	err := s.repo.Update(ctx, user)
	return err
}

func (s *Service) Lock(ctx context.Context, id string) error {
	if err := s.IsAllowed(ctx, ActionLock); err != nil {
		return err
	}
	err := s.repo.Lock(ctx, id)
	return err
}

func (s *Service) FindBySample(ctx context.Context, user *types.User) ([]*types.User, error) {
	if err := s.IsAllowed(ctx, ActionReadList); err != nil {
		return nil, err
	}
	users, err := s.repo.FindBySample(ctx, user)
	return users, err
}

func (s *Service) FindAll(ctx context.Context) ([]*types.User, error) {
	if err := s.IsAllowed(ctx, ActionReadList); err != nil {
		return nil, err
	}
	users, err := s.repo.FindAll(ctx)
	return users, err
}

func (s *Service) IsAllowed(ctx context.Context, action string) error {
	user := ctx.Value("user").(*types.User)
	isAllowed, err := s.policy.IsAllowed(ctx, user.Roles, action, PermissionResource)
	if err != nil {
		s.logger.Errorc(ctx, "failed to check permission, err: %v", err)
		return errors.Wrap(err, "failed to check permission")
	}
	if !isAllowed {
		s.logger.Errorc(ctx, "permission denied: action %s", action)
		return ErrPermissionDenied
	}
	return nil
}
