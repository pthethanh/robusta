package tutorial

import (
	"context"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"

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

// FindAll return all tutorials
func (r *MongoRepository) FindAll(ctx context.Context, offset, limit int) ([]*Tutorial, error) {
	s := r.session.Clone()
	defer s.Close()
	var tutorials []*Tutorial
	if err := r.collection(s).Find(bson.M{"status": "public"}).Sort("-created_at").Skip(offset).Limit(limit).All(&tutorials); err != nil {
		return nil, errors.Wrap(err, "failed to find all tutorials from database")
	}
	return tutorials, nil
}

// FindByID return tutorial base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*Tutorial, error) {
	s := r.session.Clone()
	defer s.Close()
	var tutorial *Tutorial
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&tutorial); err != nil {
		return nil, errors.Wrap(err, "failed to find all tutorials from database")
	}
	return tutorial, nil
}

// Create create new tutorial
func (r *MongoRepository) Create(ctx context.Context, a *Tutorial) error {
	s := r.session.Clone()
	defer s.Close()
	a.ID = db.NewID()
	a.CreatedAt = timeutil.Now()
	a.UpdatedAt = a.CreatedAt
	if err := r.collection(s).Insert(a); err != nil {
		return errors.Wrapf(err, "failed to insert tutorial %s", a.ID)
	}
	return nil
}

// Delete delete the given tutorial from database
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	if err := r.collection(s).Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"status": StatusDeleted}}); err != nil {
		return errors.Wrapf(err, "failed to delete tutorial %s", id)
	}
	return nil
}

// Update update existing tutorial
func (r *MongoRepository) Update(ctx context.Context, id string, a *Tutorial) error {
	s := r.session.Clone()
	defer s.Close()
	a.UpdatedAt = timeutil.Now()
	if err := r.collection(s).Update(bson.D{{Name: "_id", Value: id}}, a); err != nil {
		return errors.Wrapf(err, "failed to update tutorial %s", id)
	}
	return nil
}

func (r *MongoRepository) FindByCreatedTime(ctx context.Context, from time.Time, to time.Time) ([]*Tutorial, error) {
	s := r.session.Clone()
	defer s.Close()
	var tutorials []*Tutorial
	if err := r.collection(s).Find(bson.M{
		"created_at": bson.M{
			"$gt": from,
			"$lt": to,
		},
	}).All(&tutorials); err != nil {
		return nil, err
	}
	return tutorials, nil
}

func (r *MongoRepository) FindByCreatedByID(ctx context.Context, id string) ([]*Tutorial, error) {
	s := r.session.Clone()
	defer s.Close()
	var tutorials []*Tutorial
	if err := r.collection(s).Find(bson.M{"created_by_id": id}).Sort("-created_at").All(&tutorials); err != nil {
		return nil, err
	}
	return tutorials, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("goway").C("tutorial")
}

func (r *MongoRepository) Increase(ctx context.Context, id string, field string, val interface{}) error {
	s := r.session.Clone()
	defer s.Close()

	selector := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$inc": bson.M{field: val},
	}
	return r.collection(s).Update(selector, update)
}
