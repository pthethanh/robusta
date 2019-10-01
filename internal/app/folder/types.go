package folder

import "time"

type (
	Type          string
	CreateRequest struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Description string   `json:"description" validate:"min=5,max=256"`
		Type        Type     `json:"type" validate:"required"`
		Children    []string `json:"children"`
		IsPublic    bool     `json:"is_public"`

		CreatedByID     string `json:"created_by_id"`
		CreatedByName   string `json:"created_by_name"`
		CreatedByAvatar string `json:"created_by_avatar"`

		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"update_at"`
	}

	Folder struct {
		ID          string   `json:"id" bson:"_id"`
		Name        string   `json:"name" bson:"name" validate:"min=3"`
		Description string   `json:"description" bson:"description" validate:"min=5,max=256"`
		Type        Type     `json:"type" bson:"type" validate:"required"`
		Children    []string `json:"children" bson:"children"`

		CreatedByID     string `json:"created_by_id" bson:"created_by_id"`
		CreatedByName   string `json:"created_by_name" bson:"created_by_name"`
		CreatedByAvatar string `json:"created_by_avatar" bson:"created_by_avatar"`

		CreatedAt *time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt *time.Time `json:"update_at" bson:"updated_at"`
	}

	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		CreatedByID string   `json:"created_by_id"`
		Type        Type     `json:"type"`
		SortBy      []string `json:"sort_by"`
	}

	AddChildrenRequest struct {
		ID       string   `json:"id" validate:"required"`
		Children []string `json:"children" validate:"required"`
	}

	UpdateRequest struct {
		ID          string
		Name        string   `json:"name" validate:"min=3"`
		Description string   `json:"description" validate:"min=5,max=256"`
		Type        Type     `json:"type" validate:"required"`
		Children    []string `json:"children"`
	}
)

const (
	TypeChallenges Type = "challenge"
)
