package playground

import (
	"context"

	"github.com/pthethanh/robusta/internal/pkg/playground"
)

type (
	Service struct {
		runner playground.Runner
	}
)

func NewService(runner playground.Runner) *Service {
	return &Service{
		runner: runner,
	}
}

func (s *Service) Run(ctx context.Context, r *Request) (*Response, error) {
	res, err := s.runner.Run(ctx, &playground.Request{
		Code: r.Code,
	})
	if err != nil {
		return nil, err
	}
	v := Response(*res)
	return &v, nil
}
