package db

import "github.com/globalsign/mgo"

// IsErrNotFound return true if the given error is a not found error
func IsErrNotFound(err error) bool {
	return err != nil && err == mgo.ErrNotFound
}
