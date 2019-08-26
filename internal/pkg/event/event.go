package event

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
)

type (
	EventData []byte
	Event     struct {
		Type      string
		Data      EventData
		CreatedAt int64 // nanoseconds
	}

	Subscriber interface {
		Subscribe(topics ...string) <-chan Event
	}

	Publisher interface {
		Publish(event Event, topics ...string)
	}

	PublishSubscriber interface {
		Publisher
		Subscriber
	}

	Config struct {
		Buffer int `envconfig:"EVENT_BUFFER" default:"1000"`
	}
)

func NewEvent(typ string, data interface{}, createdAt time.Time) (Event, error) {
	b, err := marshal(data)
	if err != nil {
		return Event{}, err
	}
	return Event{
		Type:      typ,
		Data:      b,
		CreatedAt: createdAt.UnixNano(),
	}, nil
}

func marshal(v interface{}) (EventData, error) {
	buff := &bytes.Buffer{}
	err := json.NewEncoder(buff).Encode(v)
	return EventData(buff.Bytes()), err
}

func (data EventData) Unmarshal(v interface{}) error {
	return json.NewDecoder(bytes.NewReader(data)).Decode(v)
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}
