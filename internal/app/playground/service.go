package playground

import (
	"context"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/playground"
)

type (
	ChallengeService interface {
		Get(ctx context.Context, id string) (*types.Challenge, error)
	}

	SolutionService interface {
		Create(ctx context.Context, s *types.Solution) error
	}

	// Service is a composite service that provide ability to run code and challenges.
	Service struct {
		runner       playground.Runner
		challengeSrv ChallengeService
		solutionSrv  SolutionService
	}
)

func NewService(runner playground.Runner, challenge ChallengeService, solution SolutionService) *Service {
	return &Service{
		runner:       runner,
		challengeSrv: challenge,
		solutionSrv:  solution,
	}
}

func (s *Service) Run(ctx context.Context, r *Request) (*Response, error) {
	res, err := s.runner.Run(ctx, &playground.RunRequest{
		Code: r.Code,
	})
	if err != nil {
		return nil, err
	}
	v := Response(*res)
	return &v, nil
}

func (s *Service) Evaluate(ctx context.Context, r *EvaluateRequest) (*playground.EvaluateResponse, error) {
	challenge, err := s.challengeSrv.Get(ctx, r.ChallengeID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find challenge")
	}
	res, err := s.runner.Evaluate(ctx, &playground.EvaluateRequest{
		Solution: []byte(r.Solution),
		Test:     []byte(challenge.Test),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate solution")
	}
	status := types.SolutionStatusSuccess
	if res.IsTestFailed {
		status = types.SolutionStatusFailed
	}
	if err := s.solutionSrv.Create(ctx, &types.Solution{
		Content: r.Solution,
		Status:  status,
	}); err != nil {
		return nil, errors.Wrap(err, "failed to save solution")
	}
	return res, nil
}
