package types

const (
	EventCommentCreated = "comment_created"
	EventCommentDeleted = "comment_deleted"

	EventReactionCreated = "reaction_created"
)

type (
	ReactionChanged struct {
		IsNew       bool
		OldReaction *Reaction
		NewReaction *Reaction
		Detail      *ReactionDetail
	}
)
