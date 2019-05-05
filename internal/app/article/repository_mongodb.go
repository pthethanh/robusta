package article

import (
	"context"
	"time"

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

// FindAll return all articles
func (r *MongoRepository) FindAll(ctx context.Context, offset, limit int) ([]*types.Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var articles []*types.Article
	if err := r.collection(s).Find(nil).Sort("-created_at").Skip(offset).Limit(limit).All(&articles); err != nil {
		return nil, errors.Wrap(err, "failed to find all articles from database")
	}
	return articles, nil
}

// Get return article base on given id
func (r *MongoRepository) Get(ctx context.Context, id string) (*types.Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var article *types.Article
	if err := r.collection(s).Find(bson.M{"id": id}).One(&article); err != nil {
		return nil, errors.Wrap(err, "failed to find all articles from database")
	}
	return article, nil
}

// Create create new article
func (r *MongoRepository) Create(ctx context.Context, a *types.Article) error {
	s := r.session.Clone()
	defer s.Close()
	updatedAt := time.Now()
	a.CreatedAt = &updatedAt
	a.UpdatedAt = &updatedAt
	if err := r.collection(s).Insert(a); err != nil {
		return errors.Wrapf(err, "failed to insert article %s", a.ID)
	}
	return nil
}

// Delete delete the given article from database
func (r *MongoRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	if err := r.collection(s).Remove(bson.D{{Name: "id", Value: id}}); err != nil {
		return errors.Wrapf(err, "failed to delete article %s", id)
	}
	return nil
}

// Update update existing article
func (r *MongoRepository) Update(ctx context.Context, id string, a *types.Article) error {
	s := r.session.Clone()
	defer s.Close()
	updatedAt := time.Now()
	a.UpdatedAt = &updatedAt
	if err := r.collection(s).Update(bson.D{{Name: "id", Value: id}}, a); err != nil {
		return errors.Wrapf(err, "failed to update article %s", id)
	}
	return nil
}

func (r *MongoRepository) FindByCreatedTime(ctx context.Context, from time.Time, to time.Time) ([]*types.Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var articles []*types.Article
	if err := r.collection(s).Find(bson.M{
		"created_at": bson.M{
			"$gt": from,
			"$lt": to,
		},
	}).All(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *MongoRepository) FindByCreatedByID(ctx context.Context, id string) ([]*types.Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var articles []*types.Article
	if err := r.collection(s).Find(bson.M{"created_by_id": id}).Sort("-created_at").All(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("goway").C("article")
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
