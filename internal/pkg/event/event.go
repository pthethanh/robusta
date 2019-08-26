package event

import (
	"bytes"
	"encoding/gob"

	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
)

type (
	EventData []byte
	Event     struct {
		Type      string
		Data      EventData
		CreatedAt int // nanoseconds
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

func NewEventData(v interface{}) (EventData, error) {
	buff := &bytes.Buffer{}
	err := gob.NewEncoder(buff).Encode(v)
	return EventData(buff.Bytes()), err
}

func (data EventData) Unmarshal(v interface{}) error {
	return gob.NewDecoder(bytes.NewReader(data)).Decode(v)
}

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}
