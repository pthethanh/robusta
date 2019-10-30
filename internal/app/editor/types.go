package editor

import (
	"time"

	"github.com/pthethanh/robusta/internal/app/status"
)

type (
	LinkStatus int
	Link       struct {
		status.Status
		Success LinkStatus `json:"success"`
		Meta    LinkMeta   `json:"meta"`
	}

	LinkMeta struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Image       LinkImage `json:"image"`
	}

	LinkImage struct {
		URL string `json:"url"`
	}

	ImageToolStatus   int
	ImageToolResponse struct {
		status.Status
		Success ImageToolStatus `json:"success"`
		File    ImageToolFile   `json:"file"`
	}

	ImageToolFile struct {
		URL string `json:"url"`
	}

	Config struct {
		MaxUploadSize int64         `envconfig:"IMAGE_UPLOAD_MAX_SIZE" default:"2000000"`
		Timeout       time.Duration `envconfig:"IMAGE_UPLOAD_TIMEOUT" default:"30s"`
	}
)

const (
	LinkStatusSuccess                      = 1
	LinkStatusFailed                       = 0
	ImageToolStatusSuccess ImageToolStatus = 1
	ImageToolStatusFailed  ImageToolStatus = 0
)
