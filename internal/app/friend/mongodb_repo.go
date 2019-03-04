package friend

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/types"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
)

// MongoRepository is MongoDB implementation of repository
type MongoRepository struct {
	session *mgo.Session
}

// NewMongoRepository return new MongoDB repository
func NewMongoRepository(s *mgo.Session) *MongoRepository {
	return &MongoRepository{
		session: s,
	}
}

// FindByID return friend base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*types.Friend, error) {
	s := r.session.Clone()
	defer s.Close()
	var friend *types.Friend
	if err := r.collection(s).Find(bson.M{"id": id}).One(&friend); err != nil {
		return nil, errors.Wrap(err, "failed to find the given friend from database")
	}
	return friend, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("friends")
}
