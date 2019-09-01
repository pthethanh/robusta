package types

// Events
const (
	EventCommentCreated = "comment_created"
	EventCommentDeleted = "comment_deleted"

	EventReactionCreated = "reaction_created"

	EventNotificationArticleCommentCreated  = "notification_article_comment_created"
	EventNotificationArticleReactionCreated = "notification_article_reaction_created"

	EventNotificationCommentReactionCreated = "notification_comment_reaction_created"
	EventNotificationCommentReplyCreated    = "notification_comment_reply_created"
)

type (
	ReactionChanged struct {
		IsNew       bool
		OldReaction *Reaction
		NewReaction *Reaction
		Detail      *ReactionDetail
	}
)
