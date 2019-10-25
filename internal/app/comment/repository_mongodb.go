package comment

import (
	"context"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"
)

const (
	commentCollection = "comments"
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

func (r *MongoRepository) Create(ctx context.Context, a *types.Comment) error {
	s := r.session.Clone()
	defer s.Close()
	a.ID = db.NewID()
	a.CreatedAt = timeutil.Now()
	a.UpdatedAt = a.CreatedAt
	if err := s.DB("").C(commentCollection).Insert(a); err != nil {
		return err
	}
	return nil
}

// FindAll return all comments
func (r *MongoRepository) FindAll(ctx context.Context, req FindRequest) ([]*types.Comment, error) {
	s := r.session.Clone()
	defer s.Close()
	var comments []*types.Comment
	selector := bson.M{}
	if req.Target != "" {
		selector["target"] = req.Target
	}
	if req.ThreadID != "" {
		selector["thread_id"] = req.ThreadID
	}
	if req.ReplyToID != "" {
		selector["reply_to_id"] = req.ReplyToID
	}
	if req.CreatedByID != "" {
		selector["created_by_id"] = req.CreatedByID
	}
	sorters := req.SortBy
	if len(sorters) == 0 {
		sorters = []string{"-created_at"}
	}
	if err := s.DB("").C(commentCollection).Find(selector).Sort(sorters...).Skip(req.Offset).Limit(req.Limit).All(&comments); err != nil {
		return nil, err
	}
	return comments, nil
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
	if err := s.DB("").C(commentCollection).Update(bson.M{"_id": id}, updates); err != nil {
		return err
	}
	return nil
}

func (r *MongoRepository) Update(ctx context.Context, id string, a *types.Comment) error {
	s := r.session.Clone()
	defer s.Close()
	a.UpdatedAt = timeutil.Now()
	if err := s.DB("").C(commentCollection).Update(bson.D{{Name: "_id", Value: id}}, a); err != nil {
		return err
	}
	return nil
}

// Delete delete and return deleted comment
func (r *MongoRepository) Delete(ctx context.Context, id string) (types.Comment, error) {
	s := r.session.Clone()
	defer s.Close()
	var c types.Comment
	if _, err := s.DB("").C(commentCollection).Find(bson.M{"_id": id}).Apply(mgo.Change{
		Remove: true,
	}, &c); err != nil {
		return types.Comment{}, err
	}
	return c, nil
}

func (r *MongoRepository) FindByID(ctx context.Context, id string) (types.Comment, error) {
	s := r.session.Clone()
	defer s.Close()
	var c types.Comment
	if err := s.DB("").C(commentCollection).Find(bson.M{"_id": id}).One(&c); err != nil {
		return types.Comment{}, err
	}
	return c, nil
}
