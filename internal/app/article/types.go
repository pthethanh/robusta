package article

import (
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
)

type (
	FindRequest struct {
		Offset      int      `json:"offset" validate:"gte=0"`
		Limit       int      `json:"limit" validate:"gt=0"`
		Status      Status   `json:"-"` // don't allow set by users
		Tags        []string `json:"tags"`
		CreatedByID string   `json:"created_by_id"`
		SortBy      []string `json:"sort_by"`
	}
)

type (
	Status      string
	ContentType string
	Article     struct {
		ID          string                `json:"id" bson:"_id"`
		ArticleID   string                `json:"article_id" bson:"article_id"`
		Title       string                `json:"title" bson:"title" validate:"required"`
		Content     types.EditorJSContent `json:"content" bson:"content" validate:"required"`
		ContentType ContentType           `json:"content_type" bson:"content_type"`
		Abstract    string                `json:"abstract" bson:"abstract"`
		Source      string                `json:"source" bson:"source"`
		Status      Status                `json:"status" bson:"status"`
		PublishDate *time.Time            `json:"publish_date" bson:"publish_date"`
		Tags        []string              `json:"tags" bson:"tags"`

		Promoted     bool `json:"promoted" bson:"promoted"`
		PromotedRank int  `json:"promoted_rank" bson:"promoted_rank"`

		Views            int64 `json:"views" bson:"views"`
		Comments         int64 `json:"comments" bson:"comments"`
		ReactionUpvote   int64 `json:"reaction_upvote" bson:"reaction_upvote"`
		ReactionDownvote int64 `json:"reaction_downvote" bson:"reaction_downvote"`

		CreatedByID     string `json:"created_by_id" bson:"created_by_id"`
		CreatedByName   string `json:"created_by_name" bson:"created_by_name"`
		CreatedByAvatar string `json:"created_by_avatar" bson:"created_by_avatar"`

		CreatedAt *time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt *time.Time `json:"update_at" bson:"updated_at"`
	}
)

const (
	StatusPublic  Status = "public"
	StatusDraft   Status = "draft"
	StatusDeleted Status = "deleted"

	ContentTypeEditorJS ContentType = "editor_js"
)
