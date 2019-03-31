package policy

import "context"

type (
	Service struct {
	}
)

func (s *Service) IsAllowed(ctx context.Context, subjects []string, action string, resource string) (bool, error) {
	return true, nil
}
