package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/pkg/errors"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/uuid"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	Repository interface {
		Create(context.Context, *types.User) (string, error)
		Delete(context.Context, string) error
		UpdateInfo(context.Context, string, *types.User) error
		Lock(context.Context, string) error
		FindBySample(context.Context, *types.User) ([]*types.User, error)
		FindAll(context.Context) ([]*types.User, error)
		FindByUserID(ctx context.Context, id string) (*types.User, error)
		FindByEmail(ctx context.Context, email string) (*types.User, error)
	}

	Service struct {
		repo Repository
	}
)

// New return a new user service
func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Auth(ctx context.Context, email, password string) (*types.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil && !db.IsErrNotFound(err) {
		log.WithContext(ctx).Errorf("failed to check existing user by email, err: %v", err)
		return nil, errors.Wrap(err, "failed to check existing user by email")
	}
	if db.IsErrNotFound(err) {
		log.WithContext(ctx).Debugf("user not found, email: %s", email)
		return nil, types.ErrNotFound
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.WithContext(ctx).Error("invalid password")
		return nil, types.ErrUnauthorized
	}
	return user.Strip(), nil
}

func (s *Service) Register(ctx context.Context, req *types.RegisterRequest) (*types.User, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}
	existingUser, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil && !db.IsErrNotFound(err) {
		log.WithContext(ctx).Errorf("failed to check existing user by email, err: %v", err)
		return nil, errors.Wrap(err, "failed to check existing user by email")
	}
	if existingUser != nil {
		log.WithContext(ctx).Debug("email already registered")
		return nil, ErrEmailDuplicated
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate password")
	}
	user := &types.User{
		Password:  string(password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Status:    types.UserStatusActive,
		Provider:  types.ProviderLocal,
		UserID:    uuid.New(),
	}
	if _, err := s.repo.Create(ctx, user); err != nil {
		return nil, errors.Wrap(err, "failed to insert user")
	}
	return user.Strip(), nil
}

func (s *Service) Create(ctx context.Context, user *types.User) (string, error) {
	if err := validator.Validate(user); err != nil {
		return "", err
	}
	user.Status = types.UserStatusActive
	id, err := s.repo.Create(ctx, user)
	return id, err
}

func (s *Service) Delete(ctx context.Context, userID string) error {
	err := s.repo.Delete(ctx, userID)
	return err
}

func (s *Service) Update(ctx context.Context, userID string, user *types.User) error {
	err := s.repo.UpdateInfo(ctx, userID, user)
	return err
}

func (s *Service) Lock(ctx context.Context, userID string) error {
	err := s.repo.Lock(ctx, userID)
	return err
}

func (s *Service) FindBySample(ctx context.Context, user *types.User) ([]*types.User, error) {
	users, err := s.repo.FindBySample(ctx, user)
	return users, err
}

func (s *Service) FindAll(ctx context.Context) ([]*types.User, error) {
	users, err := s.repo.FindAll(ctx)
	return users, err
}

func (s *Service) FindByUserID(ctx context.Context, id string) (*types.User, error) {
	user, err := s.repo.FindByUserID(ctx, id)
	if err != nil && db.IsErrNotFound(err) {
		return nil, types.ErrNotFound
	}
	if err != nil && !db.IsErrNotFound(err) {
		return nil, err
	}
	return user, nil
}
