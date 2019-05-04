package user

import (
	"context"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/uuid"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"
	"github.com/pthethanh/robusta/internal/app/types"
)

type (
	MongoDBRepository struct {
		session *mgo.Session
	}
)

func (r *MongoDBRepository) Create(ctx context.Context, user *types.User) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	updatedAt := time.Now()
	user.ID = uuid.New()
	user.CreatedAt = &updatedAt
	user.UpdateAt = user.CreatedAt

	if err := r.collection(s).Insert(user); err != nil {
		return "", errors.Wrapf(err, "failed to insert user %s", user.ID)
	}
	return user.ID, nil
}

func (r *MongoDBRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Remove(bson.M{"_id": id})
}

func (r *MongoDBRepository) Update(ctx context.Context, user *types.User) error {
	s := r.session.Clone()
	defer s.Close()
	updatedAt := time.Now()
	user.UpdateAt = &updatedAt
	return r.collection(s).Update(bson.M{"_id": user.ID}, user)
}

func (r *MongoDBRepository) Lock(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"status":     "locked",
			"updated_at": time.Now(),
		},
	})
}

func (r *MongoDBRepository) FindBySample(ctx context.Context, user *types.User) ([]*types.User, error) {
	selector := bson.M{}
	if user.ID != "" {
		selector["_id"] = user.ID
	}
	if user.FirstName != "" {
		selector["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		selector["last_name"] = user.LastName
	}
	if len(user.Groups) != 0 {
		selector["groups"] = user.Groups
	}
	if len(user.Roles) != 0 {
		selector["roles"] = user.Roles
	}
	if user.Status != "" {
		selector["status"] = user.Status
	}
	if user.Username != "" {
		selector["username"] = user.Username
	}
	s := r.session.Clone()
	defer s.Close()
	var users []*types.User
	if err := r.collection(s).Find(selector).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MongoDBRepository) FindAll(context.Context, *types.User) ([]*types.User, error) {
	s := r.session.Clone()
	defer s.Close()
	var users []*types.User
	if err := r.collection(s).Find(nil).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MongoDBRepository) FindByID(ctx context.Context, id string) (*types.User, error) {
	selector := bson.M{"_id": id}
	s := r.session.Clone()
	defer s.Close()
	var user *types.User
	if err := r.collection(s).Find(selector).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MongoDBRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("goway").C("user")
}
