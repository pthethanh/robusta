package notification

import (
	"errors"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
)

var (
	ErrTimeout      = errors.New("timeout")
	ErrNotSupported = types.NewAppError(http.StatusInternalServerError, errors.New("not supported"))
)
