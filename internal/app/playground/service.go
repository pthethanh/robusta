package playground

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/lint"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/playground"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	ChallengeService interface {
		FindFolderChallengeByID(ctx context.Context, id string, folderID string) (*types.Challenge, error)
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
	return res, nil
}

func (s *Service) Evaluate(ctx context.Context, r *EvaluateRequest) (*playground.EvaluateResponse, error) {
	if err := validator.Validate(r); err != nil {
		log.WithContext(ctx).Errorf("validation failed, err: %v", err)
		return nil, status.Gen().BadRequest
	}
	challenge, err := s.challengeSrv.FindFolderChallengeByID(ctx, r.ChallengeID, r.FolderID)
	if err != nil {
		return nil, fmt.Errorf("failed to find challenge: %w", err)
	}
	switch challenge.Type {
	case types.ChallengeTypeGoExercise, "":
		return s.evaluateGoExerciseChallenge(ctx, challenge, r)
	default:
		return nil, status.Challenge().NotSupported
	}
}

func (s *Service) evaluateGoExerciseChallenge(ctx context.Context, challenge *types.Challenge, r *EvaluateRequest) (*playground.EvaluateResponse, error) {
	res, err := s.runner.Evaluate(ctx, &playground.EvaluateRequest{
		Solution: []byte(r.Solution),
		Test:     []byte(challenge.Test),
	})
	if err != nil {
		return nil, fmt.Errorf("compile failed: %w", err)
	}
	status := types.SolutionStatusSuccess
	if !res.IsSuccess() {
		status = types.SolutionStatusFailed
	}
	res.Problems = filterImportantProblems(res.Problems)
	v, err := json.Marshal(res)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to marshal evaluate result, err: %v", err)
		v = []byte(err.Error())
	}
	res.Events = nil // remove events
	if err := s.solutionSrv.Create(ctx, &types.Solution{
		Content:        r.Solution,
		Status:         status,
		EvaluateResult: string(v),
		ChallengeID:    r.ChallengeID,
	}); err != nil {
		return nil, fmt.Errorf("failed to save solution: %w", err)
	}
	return res, nil
}

func filterImportantProblems(problems []lint.Problem) []lint.Problem {
	res := make([]lint.Problem, 0)
	ignoreCategories := []string{"comments"}
	for i := 0; i < len(problems); i++ {
		ignore := false
		for _, c := range ignoreCategories {
			if problems[i].Category == c {
				ignore = true
				break
			}
		}
		if ignore {
			continue
		}
		res = append(res, problems[i])
	}
	return res
}
