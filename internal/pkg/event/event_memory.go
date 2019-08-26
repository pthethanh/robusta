package event

type (
	subscription struct {
		topics []string
		ch     chan Event
	}

	// MemoryEventStore is an event store base on memory
	MemoryEventStore struct {
		buffer        int
		subscriptions []subscription
	}
)

// NewMemoryEventStore return new memory store
func NewMemoryEventStore(conf Config) *MemoryEventStore {
	return &MemoryEventStore{
		buffer:        conf.Buffer,
		subscriptions: make([]subscription, 0),
	}
}

// Subscribe subscribe to the topics
func (pubsub *MemoryEventStore) Subscribe(topics ...string) <-chan Event {
	sub := subscription{
		topics: topics,
		ch:     make(chan Event, pubsub.buffer),
	}
	pubsub.subscriptions = append(pubsub.subscriptions, sub)
	return sub.ch
}

// Publish publish the event to the target topics
func (pubsub *MemoryEventStore) Publish(event Event, topics ...string) {
	for _, sub := range pubsub.subscriptions {
		for _, topic := range sub.topics {
			for _, pubTopic := range topics {
				if topic == pubTopic {
					sub.ch <- event
				}
			}
		}
	}
}

// Close close the underlying channels
func (pubsub *MemoryEventStore) Close() error {
	for _, sub := range pubsub.subscriptions {
		close(sub.ch)
	}
	return nil
}
