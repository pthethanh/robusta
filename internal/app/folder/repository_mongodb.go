package folder

import (
	"context"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"
)

const (
	folderCollectionName = "folders"
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

func (r *MongoDBRepository) Insert(ctx context.Context, f *Folder) error {
	s := r.session.Clone()
	defer s.Close()
	f.ID = db.NewID()
	f.CreatedAt = timeutil.Now()
	if err := s.DB("").C(folderCollectionName).Insert(f); err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) FindByID(ctx context.Context, id string) (*Folder, error) {
	s := r.session.Clone()
	defer s.Close()
	var f Folder
	if err := s.DB("").C(folderCollectionName).Find(bson.M{"_id": id}).One(&f); err != nil {
		return nil, err
	}
	return &f, nil
}

func (r *MongoDBRepository) FindAll(ctx context.Context, req FindRequest) ([]*Folder, error) {
	m := bson.M{}
	if req.CreatedByID != "" {
		m["created_by_id"] = req.CreatedByID
	}
	if req.Type != "" {
		m["type"] = req.Type
	}
	sortBy := req.SortBy
	if len(req.SortBy) == 0 {
		sortBy = []string{"-created_at"}
	}
	folders := make([]*Folder, 0)
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(folderCollectionName).Find(m).Sort(sortBy...).Skip(req.Offset).Limit(req.Limit).All(&folders); err != nil {
		return nil, err
	}
	return folders, nil
}

func (r *MongoDBRepository) Delete(cxt context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(folderCollectionName).Remove(bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) AddChildren(ctx context.Context, id string, children []string) error {
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(folderCollectionName).Update(bson.M{"_id": id}, bson.M{"$push": bson.M{"children": bson.M{"$each": children}}}); err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) Update(ctx context.Context, id string, folder Folder) error {
	s := r.session.Clone()
	defer s.Close()
	return s.DB("").C(folderCollectionName).Update(bson.M{"_id": id}, bson.M{"$set": bson.M{
		"name":        folder.Name,
		"description": folder.Description,
		"type":        folder.Type,
		"children":    folder.Children,
	}})
}
