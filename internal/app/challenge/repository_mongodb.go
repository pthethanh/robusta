package challenge

import (
	"context"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"
)

const (
	challengeCollectionName = "challenges"
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

func (r *MongoDBRepository) Insert(ctx context.Context, c *types.Challenge) error {
	s := r.session.Clone()
	defer s.Close()
	c.ID = db.NewID()
	c.CreatedAt = timeutil.Now()
	if err := s.DB("").C(challengeCollectionName).Insert(c); err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) FindByID(ctx context.Context, id string) (*types.Challenge, error) {
	s := r.session.Clone()
	defer s.Close()
	var c types.Challenge
	if err := s.DB("").C(challengeCollectionName).Find(bson.M{"_id": id}).One(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *MongoDBRepository) FindAll(ctx context.Context, req FindRequest) ([]*types.Challenge, error) {
	m := bson.M{}
	if req.CreatedByID != "" {
		m["created_by_id"] = req.CreatedByID
	}
	if len(req.Tags) > 0 {
		m["tags"] = bson.M{
			"$in": req.Tags,
		}
	}
	if len(req.IDs) > 0 {
		m["_id"] = bson.M{
			"$in": req.IDs,
		}
	}
	if req.Title != "" {
		m["title"] = &bson.RegEx{Pattern: req.Title, Options: "i"}
	}
	selects := bson.M{}
	for _, s := range req.Selects {
		selects[s] = 1
	}
	challenges := make([]*types.Challenge, 0)
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(challengeCollectionName).Find(m).Sort(req.SortBy...).Select(selects).Skip(req.Offset).Limit(req.Limit).All(&challenges); err != nil {
		return nil, err
	}
	return challenges, nil
}

func (r *MongoDBRepository) Delete(cxt context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(challengeCollectionName).Remove(bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) Update(cxt context.Context, req UpdateRequest) error {
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(challengeCollectionName).Update(bson.M{"_id": req.ID}, bson.M{
		"$set": bson.M{
			"title":       req.Title,
			"description": req.Description,
			"sample":      req.Sample,
			"test":        req.Test,
			"updated_at":  time.Now(),
		},
	}); err != nil {
		return err
	}
	return nil
}
