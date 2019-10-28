package challenge

import "github.com/pthethanh/robusta/internal/app/types"

type (
	// FindRequest hold information of finding challenges.
	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		Tags        []string `json:"tags"`
		CreatedByID string   `json:"created_by_id"`
		IDs         []string `json:"ids"`
		SortBy      []string `json:"sort_by"`
		FolderID    string   `json:"folder_id"`
		Title       string   `json:"title"`
		Selects     []string `json:"selects"`
	}

	UpdateRequest struct {
		ID          string                `json:"id" validate:"required"`
		Title       string                `json:"title" validate:"required"`
		Description types.EditorJSContent `json:"description" validate:"required"`
		Sample      string                `json:"sample" bson:"sample"`
		Test        string                `json:"test,omitempty" bson:"test"`
		Level       types.ChallengeLevel  `json:"level" bson:"level"`
	}
)
