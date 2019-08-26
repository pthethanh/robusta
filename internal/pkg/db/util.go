package db

import (
	"github.com/globalsign/mgo/bson"
)

// NewID return new id for database
func NewID() string {
	return bson.NewObjectId().Hex()
}
