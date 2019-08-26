package types

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type (
	NotificationData []byte
	NotificationType string
	Notification     struct {
		Type      NotificationType
		Data      NotificationData
		CreatedAt int64 // unix nano
	}
)

var (
	NotificationTypeEmail NotificationType = "email"
)

func (d NotificationData) Unmarshal(v interface{}) error {
	return json.Unmarshal(d, &v)
}

func NewNotification(typ NotificationType, data interface{}, createdAt time.Time) (Notification, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return Notification{}, errors.Wrap(err, "failed to marshal data")
	}
	return Notification{
		Type:      typ,
		Data:      b,
		CreatedAt: createdAt.UnixNano(),
	}, nil
}
