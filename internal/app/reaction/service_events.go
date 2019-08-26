package reaction

import (
	"context"
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func (s *Service) sendEvent(ctx context.Context, req *types.ReactionChanged, eventType string) {
	ev, err := event.NewEvent(eventType, req, time.Now())
	if err != nil {
		log.WithContext(ctx).Errorf("failed to publish event, err: %v", err)
		return
	}
	s.eventStore.Publish(ev, s.conf.Topic)
}
