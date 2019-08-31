package types

import "time"

type (
	// ReactionType type of reaction
	ReactionType string

	// ReactionTargetType is type of target
	ReactionTargetType string

	// Reactions is list of reactions
	Reactions []*Reaction

	// Reaction hold information of a reaction of a user to an object
	Reaction struct {
		ID         string             `json:"id,omitempty" bson:"_id"`
		Type       ReactionType       `json:"type" bson:"type" validate:"required,oneof=upvote downvote"`
		TargetID   string             `json:"target_id,omitempty" bson:"target_id" validate:"required"`
		TargetType ReactionTargetType `json:"target_type,omitempty" bson:"target_type" validate:"required,oneof=article comment"`

		CreatedByID     string `json:"created_by_id" bson:"created_by_id"`
		CreatedByName   string `json:"created_by_name" bson:"created_by_name"`
		CreatedByAvatar string `json:"created_by_avatar" bson:"created_by_avatar"`

		CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	}

	ReactionDetail struct {
		Upvote   int64     `json:"upvote"`
		Downvote int64     `json:"downvote"`
		Detail   Reactions `json:"reactions"`
	}
)

// Reaction types
const (
	ReactionTypeUpVote   ReactionType = "upvote"
	ReactionTypeDownVote ReactionType = "downvote"
)

// Reaction target types
const (
	ReactionTargetTypeArticle ReactionTargetType = "article"
	ReactionTargetTypeComment ReactionTargetType = "comment"
)

// Detail return summary of the reactions
func (reactions Reactions) Detail() *ReactionDetail {
	upvote := int64(0)
	downvote := int64(0)
	for _, r := range reactions {
		switch r.Type {
		case ReactionTypeUpVote:
			upvote++
		case ReactionTypeDownVote:
			downvote++
		}
	}
	return &ReactionDetail{
		Upvote:   upvote,
		Downvote: downvote,
		Detail:   reactions,
	}
}
