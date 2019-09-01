package comment

import (
	"context"
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func (s *Service) sendEvent(ctx context.Context, c types.Comment, eventType string) {
	ev, err := event.NewEvent(eventType, c, time.Now())
	if err != nil {
		log.WithContext(ctx).Errorf("failed to publish event, err: %v", err)
		return
	}
	s.es.Publish(ev, s.conf.Topic)
}

func (s *Service) listenEvents() {
	ch := s.es.Subscribe(s.conf.ReactionTopic)
	for i := 0; i < s.conf.EventWorkers; i++ {
		s.wait.Add(1)
		go func() {
			for ev := range ch {
				switch ev.Type {
				case types.EventReactionCreated:
					s.handleReactionEvent(ev)
				}
			}
			s.wait.Done()
		}()
	}
}

func (s *Service) handleReactionEvent(ev event.Event) {
	var reactionChanged types.ReactionChanged
	if err := ev.Data.Unmarshal(&reactionChanged); err != nil {
		log.Errorf("failed to unmarshal event data, err: %v", err)
		return
	}
	// not comment -> ignore...
	if reactionChanged.NewReaction.TargetType != types.ReactionTargetTypeComment {
		return
	}
	if reactionChanged.IsNew {
		log.Infof("new reaction is created on comment: %s", reactionChanged.NewReaction.TargetID)
	} else {
		log.Infof("new reaction is changed from %s to %s on comment: %s", reactionChanged.OldReaction.Type, reactionChanged.NewReaction.Type, reactionChanged.NewReaction.TargetID)
		if reactionChanged.OldReaction.Type == reactionChanged.NewReaction.Type {
			// nothing changed
			return
		}
	}
	if err := s.repo.UpdateReactions(context.Background(), reactionChanged.NewReaction.TargetID, reactionChanged.Detail); err != nil {
		log.Errorf("failed to update reaction of comment, err: %v", err)
	}
	// send notification to the comment's owner
	if err := s.sendReactionCreatedNotification(*reactionChanged.NewReaction); err != nil {
		log.Errorf("failed to send reaction created notification")
	}
}
