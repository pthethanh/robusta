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
		Save(ctx context.Context, s *types.Solution) error
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

func (s *Service) Save(ctx context.Context, solution *types.Solution) (*types.Solution, error) {
	if err := validator.Validate(solution); err != nil {
		return nil, errors.Wrap(err, "invalid solution")
	}
	user := auth.FromContext(ctx)
	if user != nil {
		solution.CreatedByID = user.UserID
		solution.CreatedByName = user.GetName()
		solution.CreatedByAvatar = user.AvatarURL
	}
	if err := s.repo.Save(ctx, solution); err != nil {
		log.WithContext(ctx).Errorf("failed to save solution, err: %v", err)
		return nil, errors.Wrap(err, "failed to save solution")
	}
	if err := s.policy.MakeOwner(ctx, policy.UserSubject(user.UserID), policy.SolutionObject(solution.ID)); err != nil {
		return nil, err
	}
	return solution, nil
}
