package policy

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

// IsCurrentUserAllowed is a util to check if the current user is allowed to do something
// the user context is expected to be existed in the given context
func IsCurrentUserAllowed(ctx context.Context, srv interface {
	IsAllowed(ctx context.Context, sub string, obj string, act string) bool
}, obj string, act string) error {
	sub := types.PolicyAnySubject
	user := auth.FromContext(ctx)
	if user != nil {
		sub = user.UserID
	}
	isAllowed := srv.IsAllowed(ctx, sub, obj, act)
	if !isAllowed {
		log.WithContext(ctx).WithFields(log.Fields{"sub": sub, "action": act, "obj": obj}).Errorf("the user is not authorized to do the action")
		return ErrNotAllowed
	}
	return nil
}
