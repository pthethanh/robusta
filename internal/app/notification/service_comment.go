package notification

import (
	"fmt"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func (s *Service) handleCommentReplyCreated(ev event.Event) {
	var n types.CommentReplyNotification
	if err := ev.Data.Unmarshal(&n); err != nil {
		log.Errorf("failed to unmarshal notification, err: %v", err)
		return
	}
	subject := fmt.Sprintf("%s has replied to your comment", n.Reply.CreatedByName)
	s.sendEmailNotification(subject, "comment_reply_created.html", n, n.Comment.CreatedByID)
}

func (s *Service) handleCommentReactionCreated(ev event.Event) {
	var n types.CommentReactionNotification
	if err := ev.Data.Unmarshal(&n); err != nil {
		log.Errorf("failed to unmarshal notification, err: %v", err)
		return
	}
	subject := fmt.Sprintf("%s %s your comment", n.Reaction.CreatedByName, n.Reaction.Type)
	s.sendEmailNotification(subject, "comment_reaction_created.html", n, n.Comment.CreatedByID)
}
