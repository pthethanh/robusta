package notification

import (
	"fmt"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func (s *Service) handleArticleCommentCreated(ev event.Event) {
	var n types.ArticleCommentNotification
	if err := ev.Data.Unmarshal(&n); err != nil {
		log.Errorf("failed to unmarshal notification, err: %v", err)
		return
	}
	subject := fmt.Sprintf("%s commented on your post: %s", n.Comment.CreatedByName, n.Article.Title)
	s.sendEmailNotification(subject, "article_comment_created.html", n, n.Article.CreatedByID)
}

func (s *Service) handleArticleReactionCreated(ev event.Event) {
	var n types.ArticleReactionNotification
	if err := ev.Data.Unmarshal(&n); err != nil {
		log.Errorf("failed to unmarshal notification, err: %v", err)
		return
	}
	subject := fmt.Sprintf("%s %s your post: %s", n.Reaction.CreatedByName, n.Reaction.Type, n.Article.Title)
	s.sendEmailNotification(subject, "article_reaction_created.html", n, n.Article.CreatedByID)
}
