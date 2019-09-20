package solution

import (
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
)

type (
	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		ChallengeID string   `json:"challenge_id"`
		CreatedByID string   `json:"created_by_id"`
		SortBy      []string `json:"sort_by"`
	}

	SolutionInfo struct {
		ID              string               `json:"id"`
		Status          types.SolutionStatus `json:"status" validate:"required"`
		CreatedAt       *time.Time           `json:"created_at"`
		CreatedByID     string               `json:"created_by_id,omitempty"`
		CreatedByName   string               `json:"created_by_name,omitempty"`
		CreatedByAvatar string               `json:"created_by_avatar,omitempty"`
	}
)
