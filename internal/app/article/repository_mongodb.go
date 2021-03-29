package article

import (
	"context"
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
func (r *MongoRepository) FindAll(ctx context.Context, req FindRequest) ([]*Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var articles []*Article
	selector := bson.M{
		"status": req.Status,
	}
	if len(req.Tags) > 0 {
		selector["tags"] = bson.M{
			"$in": req.Tags,
		}
	}
	if req.CreatedByID != "" {
		selector["created_by_id"] = req.CreatedByID
	}
	sorters := req.SortBy
	if len(sorters) == 0 {
		sorters = []string{"-created_at"}
	}
	if err := r.collection(s).Find(selector).Sort(sorters...).Skip(req.Offset).Limit(req.Limit).All(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// FindByID return article base on given id
func (r *MongoRepository) FindByID(ctx context.Context, id string) (*Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var article *Article
	if err := r.collection(s).Find(bson.M{"_id": id}).One(&article); err != nil {
		return nil, err
	}
	return article, nil
}

// Create create new article
func (r *MongoRepository) Create(ctx context.Context, a *Article) error {
	s := r.session.Clone()
	defer s.Close()
	a.ID = db.NewID()
	a.CreatedAt = timeutil.Now()
	a.UpdatedAt = a.CreatedAt
	if err := r.collection(s).Insert(a); err != nil {
		return err
	}
	return nil
}

// ChangeStatus delete the given article from database
func (r *MongoRepository) ChangeStatus(ctx context.Context, id string, status Status) error {
	s := r.session.Clone()
	defer s.Close()
	if err := r.collection(s).Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}}); err != nil {
		return err
	}
	return nil
}

// Update update existing article
func (r *MongoRepository) Update(ctx context.Context, id string, a *Article) error {
	s := r.session.Clone()
	defer s.Close()
	a.UpdatedAt = timeutil.Now()
	if err := r.collection(s).Update(bson.D{{Name: "_id", Value: id}}, a); err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) FindByCreatedTime(ctx context.Context, from time.Time, to time.Time) ([]*Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var articles []*Article
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

func (r *MongoRepository) FindByCreatedByID(ctx context.Context, id string) ([]*Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var articles []*Article
	if err := r.collection(s).Find(bson.M{"created_by_id": id}).Sort("-created_at").All(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *MongoRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("article")
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

// FindByArticleID return article base on given id
func (r *MongoRepository) FindByArticleID(ctx context.Context, id string) (*Article, error) {
	s := r.session.Clone()
	defer s.Close()
	var article *Article
	if err := r.collection(s).Find(bson.M{"article_id": id}).One(&article); err != nil {
		return nil, err
	}
	return article, nil
}

func (r *MongoRepository) UpdateReactions(ctx context.Context, id string, req *types.ReactionDetail) error {
	s := r.session.Clone()
	defer s.Close()
	updates := bson.M{
		"$set": bson.M{
			"reaction_upvote":   req.Upvote,
			"reaction_downvote": req.Downvote,
		},
	}
	if err := r.collection(s).Update(bson.M{"_id": id}, updates); err != nil {
		return err
	}
	return nil
}
