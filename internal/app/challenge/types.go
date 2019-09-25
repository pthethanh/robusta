package challenge

import "github.com/pthethanh/robusta/internal/app/types"

type (
	// FindRequest hold information of finding challenges.
	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		Tags        []string `json:"tags" validate:"required_without=CreatedByID IDs"`
		CreatedByID string   `json:"created_by_id" validate:"required_without=IDs Tags"`
		IDs         []string `json:"ids" validate:"required_without=CreatedByID Tags"`
		SortBy      []string `json:"sort_by"`
		FolderID    string   `json:"folder_id" validate:"required"`
	}

	UpdateRequest struct {
		ID          string                `json:"id" validate:"required"`
		Title       string                `json:"title" validate:"required"`
		Description types.EditorJSContent `json:"description" validate:"required"`
		Tips        string                `json:"tips" validate:"required"`
		Sample      string                `json:"sample" bson:"sample"`
		Test        string                `json:"test,omitempty" bson:"test"`
		Level       types.ChallengeLevel  `json:"level" bson:"level"`
	}
)
