package folder

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

func (s *Service) Create(ctx context.Context, req *CreateRequest) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	if err := s.isAllowed(ctx, PolicyObject, ActionCreate); err != nil {
		return err
	}
	f := &Folder{
		ID:              req.ID,
		Name:            req.Name,
		Description:     req.Description,
		Type:            req.Type,
		Children:        req.Children,
		CreatedByID:     req.CreatedByID,
		CreatedByName:   req.CreatedByName,
		CreatedByAvatar: req.CreatedByAvatar,
		CreatedAt:       req.CreatedAt,
		UpdatedAt:       req.UpdatedAt,
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
	if err := s.policy.AddPolicy(ctx, user.UserID, f.ID, types.PolicyAnyAction, types.PolicyEffectAllow); err != nil {
		return errors.Wrap(err, "failed to set permission")
	}
	if req.IsPublic {
		// make everyone permission to read this folder.
		if err := s.policy.AddPolicy(ctx, types.PolicyAnySubject, f.ID, ActionRead, types.PolicyEffectAllow); err != nil {
			return errors.Wrap(err, "failed to make the folder public read")
		}
	}
	return nil
}

func (s *Service) Get(ctx context.Context, id string) (*Folder, error) {
	f, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find the folder")
	}
	if err := s.isAllowed(ctx, id, ActionRead); err != nil {
		return nil, err
	}
	return f, nil
}

func (s *Service) FindAll(ctx context.Context, r FindRequest) ([]*Folder, error) {
	folders, err := s.repo.FindAll(ctx, r)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find folder")
	}
	rs := make([]*Folder, 0)
	for _, f := range folders {
		if err := s.isAllowed(ctx, f.ID, ActionRead); err != nil {
			log.WithContext(ctx).Debugf("user doesn't have permission on foldder %s, err: %v", f.ID, err)
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
