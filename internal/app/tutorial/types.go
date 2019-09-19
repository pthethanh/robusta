package tutorial

import "time"

type (
	Status      string
	ContentType string
	Block       struct {
		Type string                 `json:"type" bson:"type" validate:"required"`
		Data map[string]interface{} `json:"data" bson:"data"`
	}
	Content struct {
		Time    int64   `json:"time,omitempty" bson:"time" validate:"required"`
		Blocks  []Block `json:"blocks,omitempty" bson:"blocks" validate:"required,gte=1"`
		Version string  `json:"version,omitempty" bson:"version" validate:"required"`
	}
	TutorialStep struct {
		Title       string      `json:"title" bson:"title"`
		Content     Content     `json:"content" bson:"content"`
		ContentType ContentType `json:"content_type" bson:"content_type"`
	}

	Tutorial struct {
		ID        string         `json:"id" bson:"_id"`
		Title     string         `json:"title" bson:"title"`
		Steps     []TutorialStep `json:"steps"`
		Status    Status         `json:"status" bson:"status"`
		Exercises []Exercise     `json:"exercises" bson:"exercises"`

		Views    int64 `json:"views" bson:"views"`
		Likes    int64 `json:"likes" bson:"likes"`
		Comments int64 `json:"comments" bson:"comments"`

		CreatedByID     string `json:"created_by_id" bson:"created_by_id"`
		CreatedByName   string `json:"created_by_name" bson:"created_by_name"`
		CreatedByAvatar string `json:"created_by_avatar" bson:"created_by_avatar"`

		CreatedAt *time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt *time.Time `json:"update_at" bson:"updated_at"`
	}

	Exercise struct {
		ID          string `json:"id" bson:"_id"`
		Title       string `json:"title" bson:"title"`
		Description string `json:"description" bson:"description"`
		Test        string `json:"test" bson:"test"`
	}
)

const (
	StatusPublished Status = "public"
	StatusDraft     Status = "draft"
	StatusDeleted   Status = "deleted"

	ContentTypeEditorJS ContentType = "editor_js"
)

// Actions policy
const (
	ActionCreate = "tutorial:create"
	ActionDelete = "tutorial:delete"
	ActionUpdate = "tutorial:update"
	ActionRead   = "tutorial:read"
)

// Policy object name
const (
	PolicyObject = "tutorial"
)
