package folder

import (
	"context"

	"github.com/pkg/errors"
	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/policy"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	PolicyService interface {
		IsAllowed(ctx context.Context, sub string, obj string, act string) bool
		MakeOwner(ctx context.Context, sub string, obj string) error
		AddPolicy(ctx context.Context, sub string, obj string, act string, eft string) error
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
	if err := s.isAllowed(ctx, PolicyObject, ActionCreate); err != nil {
		return err
	}
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
	if err := s.policy.MakeOwner(ctx, user.UserID, f.ID); err != nil {
		return errors.Wrap(err, "failed to set permission")
	}
	if f.IsPublic {
		// it's public, grant read permission for everyone.
		if err := s.policy.AddPolicy(ctx, user.UserID, f.ID, ActionRead, policy.EffectAllow); err != nil {
			return errors.Wrap(err, "failed to grant read permission")
		}
	}
	return nil
}

func (s *Service) Get(ctx context.Context, id string) (*Folder, error) {
	if err := s.isAllowed(ctx, id, ActionRead); err != nil {
		return nil, err
	}
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
	// TODO might need to benchmark performance of checking permission.
	rs := make([]*Folder, 0)
	for _, f := range folders {
		if err := s.isAllowed(ctx, f.ID, ActionRead); err != nil {
			// user doesn't have read permission on this one, ignore.
			continue
		}
		rs = append(rs, f)
	}
	return rs, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.isAllowed(ctx, id, ActionDelete); err != nil {
		log.WithContext(ctx).Errorf("cannot delete folder, err: %v", err)
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *Service) isAllowed(ctx context.Context, id string, action string) error {
	return policy.IsCurrentUserAllowed(ctx, s.policy, id, action)
}
