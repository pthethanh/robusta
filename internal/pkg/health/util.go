package health

import (
	"context"
)

// ToCheckFunc convert the given function f to CheckFunc
func ToCheckFunc(f func() error) CheckFunc {
	return func(ctx context.Context) error {
		return f()
	}
}
