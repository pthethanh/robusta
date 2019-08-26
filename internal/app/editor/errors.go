package editor

import (
	"errors"

	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
)

// Errors that can be produced by package editor.
var (
	ErrFileSizeExceedLimit = types.NewAppError(http.StatusBadRequest, errors.New("file size exceed limit"))
)
