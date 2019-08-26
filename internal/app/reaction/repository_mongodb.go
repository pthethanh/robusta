package reaction

import (
	"context"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db"
)

const (
	collection = "reactions"
)

type (
	MongoDBRepository struct {
		session *mgo.Session
	}
)

func NewMongoDBRepository(session *mgo.Session) *MongoDBRepository {
	return &MongoDBRepository{
		session: session,
	}
}

func (r *MongoDBRepository) Upsert(ctx context.Context, reaction *types.Reaction) (bool, *types.Reaction, error) {
	s := r.session.Clone()
	defer s.Close()
	selector := bson.M{
		"target_type":   reaction.TargetType,
		"target_id":     reaction.TargetID,
		"created_by_id": reaction.CreatedByID,
	}
	updates := bson.M{
		"$set": bson.M{
			"type":            reaction.Type,
			"created_by_name": reaction.CreatedByName,
			"updated_at":      time.Now(),
		},
		"$setOnInsert": bson.M{
			"_id":           db.NewID(),
			"target_id":     reaction.TargetID,
			"target_type":   reaction.TargetType,
			"created_by_id": reaction.CreatedByID,
			"created_at":    time.Now(),
		},
	}
	change := mgo.Change{
		Update:    updates,
		Upsert:    true,
		ReturnNew: false,
	}
	var oldReaction types.Reaction
	changeInfo, err := s.DB("").C(collection).Find(selector).Apply(change, &oldReaction)
	if err != nil {
		return false, nil, err
	}
	created := changeInfo.Updated == 0
	return created, &oldReaction, nil
}

func (r *MongoDBRepository) Find(ctx context.Context, queries map[string]interface{}) (types.Reactions, error) {
	s := r.session.Clone()
	defer s.Close()
	var reactions types.Reactions
	if err := s.DB("").C(collection).Find(queries).All(&reactions); err != nil {
		return nil, err
	}
	return reactions, nil
}
