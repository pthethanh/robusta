package user

import (
	"errors"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
)

// Errors that can be thrown by package user
var (
	ErrEmailDuplicated = types.NewAppError(http.StatusBadRequest, errors.New("email already registered"))
)
