package article

import "github.com/pthethanh/robusta/internal/app/types"

type (
	FindRequest struct {
		Offset      int          `json:"offset"`
		Limit       int          `json:"limit"`
		Status      types.Status `json:"-"` // don't allow set by users
		Tags        []string     `json:"tags"`
		CreatedByID string       `json:"created_by_id"`
		SortBy      []string     `json:"sort_by"`
	}
)
