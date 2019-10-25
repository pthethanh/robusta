package user

import (
	"context"
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/util/bsonutil"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type (
	MongoDBRepository struct {
		session *mgo.Session
	}
)

func NewMongoDBRespository(session *mgo.Session) *MongoDBRepository {
	return &MongoDBRepository{
		session: session,
	}
}

func (r *MongoDBRepository) Create(ctx context.Context, user *types.User) (string, error) {
	s := r.session.Clone()
	defer s.Close()
	user.ID = db.NewID()
	user.CreatedAt = timeutil.Now()
	user.UpdateAt = user.CreatedAt

	if err := r.collection(s).Insert(user); err != nil {
		return "", err
	}
	return user.ID, nil
}

func (r *MongoDBRepository) Delete(ctx context.Context, userID string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Remove(bson.M{"user_id": userID})
}

func (r *MongoDBRepository) UpdateInfo(ctx context.Context, userID string, user *types.User) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"user_id": userID}, bson.M{
		"$set": bson.M{
			"first_name":  user.FirstName,
			"last_name":   user.LastName,
			"email":       user.Email,
			"name":        user.Name,
			"nickname":    user.NickName,
			"description": user.Description,
			"avatar_url":  user.AvatarURL,
			"location":    user.Location,
			"updated_at":  time.Now(),
		},
	})
}

func (r *MongoDBRepository) Lock(ctx context.Context, userID string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"user_id": userID}, bson.M{
		"$set": bson.M{
			"status":     types.UserStatusLocked,
			"updated_at": time.Now(),
		},
	})
}

func (r *MongoDBRepository) FindBySample(ctx context.Context, user *types.User) ([]*types.User, error) {
	selector, err := bsonutil.ToBSONMap(user)
	if err != nil {
		return nil, err
	}
	s := r.session.Clone()
	defer s.Close()
	var users []*types.User
	if err := r.collection(s).Find(selector).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MongoDBRepository) FindAll(context.Context) ([]*types.User, error) {
	s := r.session.Clone()
	defer s.Close()
	var users []*types.User
	if err := r.collection(s).Find(nil).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MongoDBRepository) FindByUserID(ctx context.Context, id string) (*types.User, error) {
	selector := bson.M{"user_id": id}
	s := r.session.Clone()
	defer s.Close()
	var user *types.User
	if err := r.collection(s).Find(selector).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MongoDBRepository) FindByEmail(ctx context.Context, email string) (*types.User, error) {
	selector := bson.M{"email": email}
	s := r.session.Clone()
	defer s.Close()
	var user *types.User
	if err := r.collection(s).Find(selector).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MongoDBRepository) UpdatePassword(ctx context.Context, userID string, newPass string) error {
	s := r.session.Clone()
	defer s.Close()
	return r.collection(s).Update(bson.M{"user_id": userID}, bson.M{
		"$set": bson.M{
			"password":   newPass,
			"updated_at": time.Now(),
		},
	})
}

func (r *MongoDBRepository) collection(s *mgo.Session) *mgo.Collection {
	return s.DB("").C("user")
}
