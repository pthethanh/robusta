package db

import (
	"errors"

	"github.com/globalsign/mgo"
)

// IsErrNotFound return true if the given error is a not found error
func IsErrNotFound(err error) bool {
	return errors.Is(err, mgo.ErrNotFound)
}
