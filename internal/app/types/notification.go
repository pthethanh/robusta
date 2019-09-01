package types

type (
	ArticleCommentNotification struct {
		Article ArticleInfo `json:"article"`
		Comment Comment     `json:"comment"`
	}

	ArticleReactionNotification struct {
		Article  ArticleInfo `json:"article"`
		Reaction Reaction    `json:"reaction"`
	}
)

type (
	CommentReplyNotification struct {
		ParentOwner string  `json:"parent_owner,omitempty"`
		Comment     Comment `json:"comment"`
	}

	CommentReactionNotification struct {
		ParentOwner string   `json:"parent_owner,omitempty"`
		Reaction    Reaction `json:"reaction"`
	}
)
