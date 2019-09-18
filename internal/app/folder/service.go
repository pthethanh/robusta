package folder

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/policy"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	PolicyService interface {
		IsAllowed(ctx context.Context, sub policy.Subject, obj policy.Object, act policy.Action) bool
		MakeOwner(ctx context.Context, sub policy.Subject, obj policy.Object) error
	}

	Repository interface {
		Insert(ctx context.Context, c *Folder) error
		FindByID(ctx context.Context, id string) (*Folder, error)
		FindAll(ctx context.Context, r FindRequest) ([]*Folder, error)
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

func (s *Service) Create(ctx context.Context, f *Folder) error {
	if err := validator.Validate(f); err != nil {
		return err
	}
	user := auth.FromContext(ctx)
	if user != nil {
		f.CreatedByID = user.UserID
		f.CreatedByName = user.GetName()
		f.CreatedByAvatar = user.AvatarURL
	}
	if err := s.repo.Insert(ctx, f); err != nil {
		return errors.Wrap(err, "failed to insert folder")
	}
	if err := s.policy.MakeOwner(ctx, policy.UserSubject(user.UserID), policy.FolderObject(f.ID)); err != nil {
		return errors.Wrap(err, "failed to set permission")
	}
	return nil
}

func (s *Service) Get(ctx context.Context, id string) (*Folder, error) {
	f, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find the folder")
	}
	return f, nil
}

func (s *Service) FindAll(ctx context.Context, r FindRequest) ([]*Folder, error) {
	folders, err := s.repo.FindAll(ctx, r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find folder")
	}
	return folders, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := policy.IsCurrentUserAllowed(ctx, s.policy, policy.FolderObject(id), policy.ActionDelete); err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
