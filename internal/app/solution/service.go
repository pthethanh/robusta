package solution

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
	Repository interface {
		Insert(ctx context.Context, s *types.Solution) error
		FindAll(ctx context.Context, req FindRequest) ([]*types.Solution, error)
	}

	PolicyService interface {
		IsAllowed(ctx context.Context, sub policy.Subject, obj policy.Object, act policy.Action) bool
		MakeOwner(ctx context.Context, sub policy.Subject, obj policy.Object) error
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

func (s *Service) Create(ctx context.Context, solution *types.Solution) error {
	if err := validator.Validate(solution); err != nil {
		return errors.Wrap(err, "invalid solution")
	}
	user := auth.FromContext(ctx)
	if user != nil {
		solution.CreatedByID = user.UserID
		solution.CreatedByName = user.GetName()
		solution.CreatedByAvatar = user.AvatarURL
	}
	if err := s.repo.Insert(ctx, solution); err != nil {
		log.WithContext(ctx).Errorf("failed to save solution, err: %v", err)
		return errors.Wrap(err, "failed to save solution")
	}
	if err := s.policy.MakeOwner(ctx, policy.UserSubject(user.UserID), policy.SolutionObject(solution.ID)); err != nil {
		return err
	}
	return nil
}

func (s *Service) FindAll(ctx context.Context, req FindRequest) ([]*types.Solution, error) {
	solutions, err := s.repo.FindAll(ctx, req)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to find solutions from database, err: %v", err)
		return nil, errors.Wrap(err, "failed to find solutions from database")
	}
	return solutions, nil
}
