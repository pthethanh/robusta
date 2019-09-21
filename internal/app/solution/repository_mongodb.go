package solution

import (
	"context"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/db"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"
)

const (
	solutionCollectionName = "solutions"
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

func (r *MongoDBRepository) Insert(ctx context.Context, solution *types.Solution) error {
	solution.ID = db.NewID()
	solution.CreatedAt = timeutil.Now()
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(solutionCollectionName).Insert(solution); err != nil {
		return err
	}
	return nil
}

func (r *MongoDBRepository) FindAll(ctx context.Context, req FindRequest) ([]*types.Solution, error) {
	m := bson.M{}
	if req.CreatedByID != "" {
		m["created_by_id"] = req.CreatedByID
	}
	if len(req.ChallengeIDs) > 0 {
		m["challenge_id"] = bson.M{
			"$in": req.ChallengeIDs,
		}
	}
	if req.Status != "" {
		m["status"] = req.Status
	}
	sortBy := req.SortBy
	if len(req.SortBy) == 0 {
		sortBy = []string{"-created_at"}
	}
	solutions := make([]*types.Solution, 0)
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(solutionCollectionName).Find(m).Sort(sortBy...).Skip(req.Offset).Limit(req.Limit).All(&solutions); err != nil {
		return nil, err
	}
	return solutions, nil
}

func (r *MongoDBRepository) FindByID(ctx context.Context, id string) (*types.Solution, error) {
	s := r.session.Clone()
	defer s.Close()
	var solution types.Solution
	if err := s.DB("").C(solutionCollectionName).Find(bson.M{"_id": id}).One(&solution); err != nil {
		return nil, err
	}
	return &solution, nil
}
