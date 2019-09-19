package policy

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

// IsCurrentUserAllowed is a util to check if the current user is allowed to do something
// the user context is expected to be existed in the given context
func IsCurrentUserAllowed(ctx context.Context, srv interface {
	IsAllowed(ctx context.Context, sub string, obj string, act string) bool
}, obj string, act string) error {
	user := auth.FromContext(ctx)
	if user == nil {
		return ErrNotAllowed
	}
	isAllowed := srv.IsAllowed(ctx, user.UserID, obj, act)
	if !isAllowed {
		log.WithContext(ctx).WithFields(log.Fields{"user_id": user.UserID, "action": act, "obj": obj}).Errorf("the user is not authorized to do the action")
		return ErrNotAllowed
	}
	return nil
}
