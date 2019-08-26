package reaction

import (
	"context"
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func (s *Service) sendEvent(ctx context.Context, req *types.ReactionChanged, eventType string) {
	eventData, err := event.NewEventData(req)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to publish event, err: %v", err)
		return
	}
	s.eventStore.Publish(event.Event{
		Type:      eventType,
		Data:      eventData,
		CreatedAt: time.Now().Nanosecond(),
	}, s.conf.Topic)
}
