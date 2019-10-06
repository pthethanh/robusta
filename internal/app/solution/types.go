package solution

import (
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
)

type (
	FindRequest struct {
		Offset        int      `json:"offset"`
		Limit         int      `json:"limit"`
		ChallengeIDs  []string `json:"challenge_ids"`
		CreatedByID   string   `json:"created_by_id"`
		Status        string   `json:"status"`
		SortBy        []string `json:"sort_by"`
		IncludeDetail bool     `json:"include_detail"`
	}

	SolutionInfo struct {
		ID              string               `json:"id"`
		ChallengeID     string               `json:"challenge_id"`
		Status          types.SolutionStatus `json:"status"`
		CreatedAt       *time.Time           `json:"created_at"`
		CreatedByID     string               `json:"created_by_id,omitempty"`
		CreatedByName   string               `json:"created_by_name,omitempty"`
		CreatedByAvatar string               `json:"created_by_avatar,omitempty"`
		Content         string               `json:"content,omitempty"`
	}

	CompletionReportRequest struct {
		ChallengeIDs  []string `json:"challenge_ids" validate:"required"`
		CreatedByID   string   `json:"created_by_id"`
		IncludeDetail bool     `json:"include_detail"`
		Status        string   `json:"status"`
	}

	solutionInfoByCreatedAt []SolutionInfo
)

func (l solutionInfoByCreatedAt) Len() int {
	return len(l)
}

func (l solutionInfoByCreatedAt) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l solutionInfoByCreatedAt) Less(i, j int) bool {
	return l[i].CreatedAt.Before(*l[j].CreatedAt)
}
