package user

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/jwt"
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
		UpdatePassword(ctx context.Context, userID string, newPassword string) error
	}

	PolicyService interface {
		Validate(ctx context.Context, obj string, act string) error
	}

	Config struct {
		ResetPasswordTokenLifeTime time.Duration `envconfig:"USER_RESET_PASSWORD_TOKEN_LIFE_TIME" default:"15m"`
		NotificationTopic          string        `envconfig:"NOTIFICATION_TOPIC" default:"r_topic_notification"`
	}

	Service struct {
		conf   Config
		repo   Repository
		policy PolicyService
		jwt    jwt.SignVerifier
		es     event.Publisher
	}
)

// New return a new user service
func New(conf Config, repo Repository, policy PolicyService, jwtSigner jwt.SignVerifier, es event.Publisher) *Service {
	return &Service{
		conf:   conf,
		repo:   repo,
		policy: policy,
		jwt:    jwtSigner,
		es:     es,
	}
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func (s *Service) Auth(ctx context.Context, email, password string) (*types.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil && !db.IsErrNotFound(err) {
		log.WithContext(ctx).Errorf("failed to check existing user by email, err: %v", err)
		return nil, status.Gen().Internal
	}
	if db.IsErrNotFound(err) {
		log.WithContext(ctx).Debugf("user not found, email: %s", email)
		return nil, status.Auth().InvalidUserPassword
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.WithContext(ctx).Error("invalid password")
		return nil, status.Auth().InvalidUserPassword
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
		return nil, fmt.Errorf("failed to check existing user by email: %w", err)
	}
	if existingUser != nil {
		log.WithContext(ctx).Debug("email already registered")
		return nil, status.User().DuplicatedEmail
	}
	password, err := s.generatePassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to generate password: %w", err)
	}
	user := &types.User{
		Password:  password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Status:    types.UserStatusActive,
		Provider:  types.ProviderLocal,
		UserID:    uuid.New(),
	}
	if _, err := s.repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
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

func (s *Service) FindAll(ctx context.Context) ([]*types.UserInfo, error) {
	if err := s.policy.Validate(ctx, types.PolicyObjectUser, types.PolicyActionUserReadList); err != nil {
		return nil, err
	}
	users, err := s.repo.FindAll(ctx)
	info := make([]*types.UserInfo, 0)
	for _, usr := range users {
		info = append(info, &types.UserInfo{
			ID:          usr.ID,
			Email:       usr.Email,
			FirstName:   usr.FirstName,
			LastName:    usr.LastName,
			Name:        usr.GetName(),
			NickName:    usr.NickName,
			Description: usr.Description,
			UserID:      usr.UserID,
			AvatarURL:   usr.AvatarURL,
		})
	}
	return info, err
}

func (s *Service) FindByUserID(ctx context.Context, id string) (*types.User, error) {
	user, err := s.repo.FindByUserID(ctx, id)
	if err != nil && db.IsErrNotFound(err) {
		return nil, status.Gen().NotFound
	}
	if err != nil && !db.IsErrNotFound(err) {
		return nil, err
	}
	return user, nil
}

func (s *Service) GenerateResetPasswordToken(ctx context.Context, mail string) (string, error) {
	user, err := s.repo.FindByEmail(ctx, mail)
	if err != nil {
		return "", err
	}
	token, err := s.jwt.Sign(jwt.Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        user.UserID,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(s.conf.ResetPasswordTokenLifeTime).Unix(),
		},
		UserID: user.UserID,
	})
	if err != nil {
		return "", err
	}
	// notify to user
	ev, err := event.NewEvent(types.EventPasswordResetTokenCreated, types.ResetPasswordTokenCreated{
		Token: token,
		User: types.UserInfo{
			ID:        user.ID,
			Email:     user.Email,
			UserID:    user.UserID,
			Name:      user.Name,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}, time.Now())
	if err != nil {
		log.WithContext(ctx).Errorf("failed to create notification event, err: %v", err)
		return "", fmt.Errorf("failed to create event: %w", err)
	}
	s.es.Publish(ev, s.conf.NotificationTopic)

	return token, nil
}

func (s *Service) ResetPassword(ctx context.Context, r ResetPasswordRequest) error {
	if err := validator.Validate(r); err != nil {
		log.WithContext(ctx).Errorf("invalid request, err: %v", err)
		return err
	}
	c, err := s.jwt.Verify(r.Token)
	if err != nil {
		log.WithContext(ctx).Errorf("invalid token, err: %v", err)
		return err
	}
	// verify user exist
	if _, err := s.repo.FindByUserID(ctx, c.UserID); err != nil {
		log.WithContext(ctx).Errorf("could not found user, err: %v", err)
		return err
	}
	pass, err := s.generatePassword(r.NewPassword)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to generate password, err: %v", err)
		return err
	}
	// update pass
	if err := s.repo.UpdatePassword(ctx, c.UserID, pass); err != nil {
		log.WithContext(ctx).Errorf("failed to update password, err: %v", err)
		return err
	}
	return nil
}

func (s *Service) generatePassword(pass string) (string, error) {
	rs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate password: %w", err)
	}
	return string(rs), nil
}
