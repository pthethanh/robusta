package comment

import (
	"context"
	"fmt"
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
)

func (s *Service) sendReactionCreatedNotification(r types.Reaction) error {
	c, err := s.FindByID(context.Background(), r.TargetID)
	if err != nil {
		return fmt.Errorf("failed to find comment: %w", err)
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
		return fmt.Errorf("failed to create comment reaction notification event: %w", err)
	}
	s.es.Publish(ev, s.conf.NotificationTopic)
	return nil
}

func (s *Service) sendReplyCreatedNotification(reply types.Comment) error {
	c, err := s.FindByID(context.Background(), reply.ReplyToID)
	if err != nil {
		return fmt.Errorf("failed to find original comment: %w", err)
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
		return fmt.Errorf("failed to create reply comment notification event: %w", err)
	}
	s.es.Publish(ev, s.conf.NotificationTopic)
	return nil
}
