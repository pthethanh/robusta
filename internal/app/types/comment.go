package types

import "time"

type (
	CommentTargetType string
	Comment           struct {
		ID         string            `json:"id,omitempty" bson:"_id"`
		Target     string            `json:"target,omitempty" bson:"target" validate:"required"`
		TargetType CommentTargetType `json:"target_type,omitempty" bson:"target_type" validate:"required"`
		Content    string            `json:"content,omitempty" bson:"content" validate:"required"`
		ReplyToID  string            `json:"reply_to_id,omitempty" bson:"reply_to_id"`
		ThreadID   string            `json:"thread_id,omitempty" bson:"thread_id"`
		Level      int               `json:"level" bson:"level"`

		ReactionUpvote   int64 `json:"reaction_upvote" bson:"reaction_upvote"`
		ReactionDownvote int64 `json:"reaction_downvote" bson:"reaction_downvote"`

		CreatedByName   string     `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByID     string     `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedByAvatar string     `json:"created_by_avatar" bson:"created_by_avatar"`
		CreatedAt       *time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt       *time.Time `json:"modified_at,omitempty" bson:"updated_at"`
	}
)

const (
	CommentTargetTypeArticle CommentTargetType = "article"
)
