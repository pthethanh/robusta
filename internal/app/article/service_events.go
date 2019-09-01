package article

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func (s *Service) handleEvents() {
	ch := s.es.Subscribe(s.conf.CommentTopic, s.conf.ReactionTopic)
	for i := 0; i < s.conf.EventWorkers; i++ {
		s.wait.Add(1)
		go func() {
			for e := range ch {
				switch e.Type {
				case types.EventCommentCreated, types.EventCommentDeleted:
					s.handleCommentEvent(e)
				case types.EventReactionCreated:
					s.handleReactionEvent(e)
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
	// not article -> ignore...
	if reactionChanged.NewReaction.TargetType != types.ReactionTargetTypeArticle {
		return
	}
	if reactionChanged.IsNew {
		log.Infof("new reaction is created on article: %s", reactionChanged.NewReaction.TargetID)
	} else {
		log.Infof("new reaction is changed from %s to %s on article: %s", reactionChanged.OldReaction.Type, reactionChanged.NewReaction.Type, reactionChanged.NewReaction.TargetID)
		if reactionChanged.OldReaction.Type == reactionChanged.NewReaction.Type {
			// nothing changed
			return
		}
	}
	if err := s.repo.UpdateReactions(context.Background(), reactionChanged.NewReaction.TargetID, reactionChanged.Detail); err != nil {
		log.Errorf("failed to update reaction of article, err: %v", err)
	}
	if err := s.sendReactionCreatedNotification(*reactionChanged.NewReaction); err != nil {
		log.Errorf("failed to send reaction created notification")
	}
}

func (s *Service) handleCommentEvent(ev event.Event) {
	log.Debugf("article: received event: %s, created_at: %v", ev.Type, ev.CreatedAt)
	var c types.Comment
	if err := ev.Data.Unmarshal(&c); err != nil {
		log.Errorf("failed to unmarshal event data, err: %v", err)
		return
	}
	// not article -> ignore...
	if c.TargetType != types.CommentTargetTypeArticle {
		return
	}
	switch ev.Type {
	case types.EventCommentCreated:
		if err := s.updateCommentStatistic(c, +1); err != nil {
			log.Errorf("failed to increase comment number of article: %s, err: %v", c.Target, err)
		}
		if err := s.sendCommentCreatedNotification(c); err != nil {
			log.Errorf("failed to send comment created notification")
		}
	case types.EventCommentDeleted:
		if err := s.updateCommentStatistic(c, -1); err != nil {
			log.Errorf("failed to decrease comment number of article: %s, err: %v", c.Target, err)
		}
	}
}

func (s *Service) updateCommentStatistic(c types.Comment, inc int) error {
	if err := s.repo.Increase(context.Background(), c.Target, "comments", inc); err != nil {
		return err
	}
	return nil
}
