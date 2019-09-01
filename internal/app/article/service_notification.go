package article

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
)

func (s *Service) sendCommentCreatedNotification(c types.Comment) error {
	a, err := s.FindByID(context.Background(), c.Target)
	if err != nil {
		return errors.Wrap(err, "failed to find article")
	}
	if a.CreatedByID == c.CreatedByID {
		return nil
	}
	ev, err := event.NewEvent(types.EventNotificationArticleCommentCreated, types.ArticleCommentNotification{
		Article: types.ArticleInfo{
			ID:              a.ID,
			Title:           a.Title,
			CreatedByID:     a.CreatedByID,
			CreatedByName:   a.CreatedByName,
			CreatedByAvatar: a.CreatedByAvatar,
		},
		Comment: c,
	}, time.Now())
	if err != nil {
		return errors.Wrap(err, "failed to create article comment notification event")
	}
	s.es.Publish(ev, s.conf.NotificationTopic)
	return nil
}

func (s *Service) sendReactionCreatedNotification(r types.Reaction) error {
	a, err := s.FindByID(context.Background(), r.TargetID)
	if err != nil {
		return errors.Wrap(err, "failed to find article")
	}
	if a.CreatedByID == r.CreatedByID {
		return nil
	}
	ev, err := event.NewEvent(types.EventNotificationArticleReactionCreated, types.ArticleReactionNotification{
		Article: types.ArticleInfo{
			ID:              a.ID,
			Title:           a.Title,
			CreatedByID:     a.CreatedByID,
			CreatedByName:   a.CreatedByName,
			CreatedByAvatar: a.CreatedByAvatar,
		},
		Reaction: r,
	}, time.Now())
	if err != nil {
		return errors.Wrap(err, "failed to create article reaction notification event")
	}
	s.es.Publish(ev, s.conf.NotificationTopic)
	return nil
}
