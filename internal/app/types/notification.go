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
		Comment Comment `json:"comment"`
		Reply   Comment `json:"reply"`
	}

	CommentReactionNotification struct {
		Comment  Comment  `json:"comment"`
		Reaction Reaction `json:"reaction"`
	}
)
