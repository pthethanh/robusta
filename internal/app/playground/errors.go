package playground

import (
	"errors"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
)

var (
	NotSupported = types.NewAppError(http.StatusBadRequest, errors.New("challenge type not supported"))
)
