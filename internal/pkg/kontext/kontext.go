package kontext

import (
	"context"

	"github.com/pthethanh/robusta/internal/pkg/uuid"
)

var (
	requestIDContextKey = struct{}{}
)

// RequestIDFromContext return request_id in the given context if one is existed
func RequestIDFromContext(ctx context.Context) string {
	if v := ctx.Value(requestIDContextKey); v != nil {
		return v.(string)
	}
	return ""
}

// NewRequestIDContext create new context with request_id inside
func NewRequestIDContext(ctx context.Context) context.Context {
	newCtx := context.WithValue(ctx, requestIDContextKey, uuid.New())
	return newCtx
}
