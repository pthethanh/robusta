package uuid

import (
	"github.com/google/uuid"
)

// New return new unique ID
func New() string {
	return uuid.New().String()
}
