package folder

import (
	"context"
	"fmt"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	PolicyService interface {
		Validate(ctx context.Context, obj string, act string) error
		AddPolicy(ctx context.Context, p types.Policy) error
	}

	Repository interface {
		Insert(ctx context.Context, c *Folder) error
		FindByID(ctx context.Context, id string) (*Folder, error)
		FindAll(ctx context.Context, r FindRequest) ([]*Folder, error)
		Delete(cxt context.Context, id string) error
		AddChildren(ctx context.Context, id string, children []string) error
		Update(ctx context.Context, id string, folder Folder) error
	}
	Config struct {
		MaxPageSize int `envconfig:"FOLDER_MAX_PAGE_SIZE" default:"50"`
	}
	Service struct {
		conf   Config
		repo   Repository
		policy PolicyService
	}
)

func NewService(conf Config, repo Repository, policy PolicyService) *Service {
	return &Service{
		repo:   repo,
		policy: policy,
	}
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func (s *Service) Create(ctx context.Context, req *CreateRequest) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	if err := s.policy.Validate(ctx, types.PolicyObjectFolder, types.PolicyActionFolderCreate); err != nil {
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
		return fmt.Errorf("failed to insert folder: %w", err)
	}
	if err := s.policy.AddPolicy(auth.NewAdminContext(ctx), types.Policy{
		Subject: user.UserID,
		Object:  f.ID,
		Action:  types.PolicyActionAny,
		Effect:  types.PolicyEffectAllow,
	}); err != nil {
		return fmt.Errorf("failed to set permission: %w", err)
	}
	if req.IsPublic {
		// make everyone permission to read this folder.
		if err := s.policy.AddPolicy(auth.NewAdminContext(ctx), types.Policy{
			Subject: types.PolicySubjectAny,
			Object:  f.ID,
			Action:  types.PolicyActionFolderRead,
			Effect:  types.PolicyEffectAllow,
		}); err != nil {
			return fmt.Errorf("failed to make the folder public read: %w", err)
		}
	}
	return nil
}

func (s *Service) Get(ctx context.Context, id string) (*Folder, error) {
	if err := s.policy.Validate(ctx, id, types.PolicyActionFolderRead); err != nil {
		return nil, err
	}
	f, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find the folder: %w", err)
	}
	return f, nil
}

func (s *Service) FindAll(ctx context.Context, r FindRequest) ([]*Folder, error) {
	folders, err := s.repo.FindAll(ctx, r)
	if err != nil {
		return nil, fmt.Errorf("failed to find folder: %w", err)
	}
	if r.Limit > s.conf.MaxPageSize {
		r.Limit = s.conf.MaxPageSize
	}
	rs := make([]*Folder, 0)
	for _, f := range folders {
		if err := s.policy.Validate(ctx, f.ID, types.PolicyActionFolderRead); err != nil {
			log.WithContext(ctx).Debugf("user doesn't have permission on foldder %s, err: %v", f.ID, err)
			continue
		}
		rs = append(rs, f)
	}
	return rs, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	if err := s.policy.Validate(ctx, id, types.PolicyActionFolderDelete); err != nil {
		log.WithContext(ctx).Errorf("cannot delete folder, err: %v", err)
		return err
	}
	return s.repo.Delete(ctx, id)
}

func (s *Service) AddChildren(ctx context.Context, req AddChildrenRequest) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	if err := s.policy.Validate(ctx, types.PolicyObjectFolder, types.PolicyActionFolderUpdate); err != nil {
		return err
	}
	if err := s.repo.AddChildren(ctx, req.ID, req.Children); err != nil {
		return fmt.Errorf("failed to add children: %w", err)
	}
	return nil
}

func (s *Service) Update(ctx context.Context, req UpdateRequest) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	if err := s.policy.Validate(ctx, types.PolicyObjectFolder, types.PolicyActionFolderUpdate); err != nil {
		return err
	}
	if err := s.repo.Update(ctx, req.ID, Folder{
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		Children:    req.Children,
	}); err != nil {
		return fmt.Errorf("failed to update folder: %w", err)
	}
	return nil
}
