package solution

import (
	"context"
	"sort"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/app/utils/policyutil"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/validator"
)

type (
	Repository interface {
		Insert(ctx context.Context, s *types.Solution) error
		FindAll(ctx context.Context, req FindRequest) ([]*types.Solution, error)
		FindByID(ctx context.Context, id string) (*types.Solution, error)
	}

	PolicyService interface {
		IsAllowed(ctx context.Context, sub string, obj string, act string) bool
		MakeOwner(ctx context.Context, sub string, obj string) error
	}

	Config struct {
		MaxPageSize int `envconfig:"SOLUTION_MAX_PAGE_SIZE" default:"50"`
	}
	Service struct {
		conf   Config
		repo   Repository
		policy PolicyService
	}
)

func NewService(conf Config, repo Repository, policy PolicyService) *Service {
	return &Service{
		conf:   conf,
		repo:   repo,
		policy: policy,
	}
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
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
	if user == nil {
		return nil
	}
	if err := s.policy.MakeOwner(ctx, user.UserID, solution.ID); err != nil {
		return err
	}
	return nil
}

// FindSolutionInfo return information of a list of solution base on the given request.
// The content and detail of solution are striped from the result, hence this method
// is safe to call without checking permission.
func (s *Service) FindSolutionInfo(ctx context.Context, req FindRequest) ([]SolutionInfo, error) {
	if req.Limit > s.conf.MaxPageSize {
		req.Limit = s.conf.MaxPageSize
	}
	user := auth.FromContext(ctx)
	isOwner := user.UserID == req.CreatedByID
	// if not the owner get his/her solutions, need to check permissions
	if req.IncludeDetail && !isOwner {
		if err := s.isAllowed(ctx, types.PolicyObjectSolution, types.PolicyActionSolutionReadListDetail); err != nil {
			log.WithContext(ctx).Errorf("failed to read solution detail, err: %v", err)
			return nil, err
		}
	}
	solutions, err := s.repo.FindAll(ctx, req)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to find solutions from database, err: %v", err)
		return nil, errors.Wrap(err, "failed to find solutions from database")
	}
	infos := make([]SolutionInfo, 0)
	for _, s := range solutions {
		info := SolutionInfo{
			ID:              s.ID,
			ChallengeID:     s.ChallengeID,
			Status:          s.Status,
			CreatedAt:       s.CreatedAt,
			CreatedByID:     s.CreatedByID,
			CreatedByName:   s.CreatedByName,
			CreatedByAvatar: s.CreatedByAvatar,
		}
		if req.IncludeDetail {
			info.Content = s.Content
		}
		infos = append(infos, info)
	}
	return infos, nil
}

func (s *Service) Get(ctx context.Context, id string) (*types.Solution, error) {
	if err := s.isAllowed(ctx, id, types.PolicyActionSolutionRead); err != nil {
		return nil, err
	}
	return s.repo.FindByID(ctx, id)
}

// GetCompletionReport return last success/failure solutions
func (s *Service) GetCompletionReport(ctx context.Context, r CompletionReportRequest) ([]SolutionInfo, error) {
	if err := validator.Validate(r); err != nil {
		log.WithContext(ctx).Errorf("invalid input, err: %v", err)
		return nil, err
	}
	offset := 0
	rs := make([]SolutionInfo, 0)
	seen := make(map[string]bool)
	for {
		solutions, err := s.FindSolutionInfo(ctx, FindRequest{
			Offset:        offset,
			Limit:         s.conf.MaxPageSize,
			ChallengeIDs:  r.ChallengeIDs,
			CreatedByID:   r.CreatedByID,
			Status:        r.Status,
			SortBy:        []string{"challenge_id", "created_by_id", "-status", "-created_at"},
			IncludeDetail: r.IncludeDetail,
		})
		if err != nil {
			log.WithContext(ctx).Errorf("failed to find solutions, err: %v", err)
			return nil, err
		}
		for _, s := range solutions {
			k := s.ChallengeID + s.CreatedByID
			if _, ok := seen[k]; !ok {
				rs = append(rs, s)
				seen[k] = true
			}
		}
		if len(solutions) < s.conf.MaxPageSize {
			// no more...
			break
		}
		offset += len(solutions)
	}
	sort.Sort(sort.Reverse(solutionInfoByCreatedAt(rs)))
	return rs, nil
}

func (s *Service) isAllowed(ctx context.Context, id string, act string) error {
	return policyutil.IsCurrentUserAllowed(ctx, s.policy, id, act)
}
