package comment

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
)

func (s *Service) sendReactionCreatedNotification(r types.Reaction) error {
	c, err := s.FindByID(context.Background(), r.TargetID)
	if err != nil {
		return errors.Wrap(err, "failed to find comment")
	}
	// don't need to send notify to the same person
	if c.CreatedByID == r.CreatedByID {
		return nil
	}
	ev, err := event.NewEvent(types.EventNotificationCommentReactionCreated, types.CommentReactionNotification{
		Comment:  c,
		Reaction: r,
	}, time.Now())
	if err != nil {
		return errors.Wrap(err, "failed to create comment reaction notification event")
	}
	s.es.Publish(ev, s.conf.NotificationTopic)
	return nil
}

func (s *Service) sendReplyCreatedNotification(reply types.Comment) error {
	c, err := s.FindByID(context.Background(), reply.ReplyToID)
	if err != nil {
		return errors.Wrap(err, "failed to find original comment")
	}
	// don't need to send notify to the same person
	if reply.CreatedByID == c.CreatedByID {
		return nil
	}
	ev, err := event.NewEvent(types.EventNotificationCommentReplyCreated, types.CommentReplyNotification{
		Comment: c,
		Reply:   reply,
	}, time.Now())
	if err != nil {
		return errors.Wrap(err, "failed to create reply comment notification event")
	}
	s.es.Publish(ev, s.conf.NotificationTopic)
	return nil
}
