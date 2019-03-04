package db

import "github.com/globalsign/mgo"

const (
	TypeMongoDB = "mongodb"
	TypeMySQL   = "mysql"
)

type (
	// Connections all supported types of database connections
	Connections struct {
		Type    string
		MongoDB *mgo.Session
	}
)

// IsErrNotFound return true if the given error is a not found error
func IsErrNotFound(err error) bool {
	return err == mgo.ErrNotFound
}

// Close close all underlying connections
func (c *Connections) Close() error {
	switch c.Type {
	case TypeMongoDB:
		if c.MongoDB != nil {
			c.MongoDB.Close()
		}
	}
	return nil
}
