package policy

import (
	"context"
	"strings"
)

type (
	Service struct {
	}
)

const (
	roleAdmin = "admin"
)

// IsAllowed check if the subject is allowed to do the action on the given resources
// TODO this is just a mock implementation, please correct me
func (s *Service) IsAllowed(ctx context.Context, roles []string, action string, resource string) (bool, error) {
	// admin is allowed to do anything he wants
	for _, role := range roles {
		if role == roleAdmin {
			return true, nil
		}
	}
	// non-admin is allowed to read only
	if strings.Contains(action, "read") {
		return false, nil
	}
	return true, nil
}
