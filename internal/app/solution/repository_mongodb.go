package solution

import (
	"context"

	"github.com/globalsign/mgo"
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

func (r *MongoDBRepository) Save(ctx context.Context, solution *types.Solution) error {
	solution.ID = db.NewID()
	solution.CreatedAt = timeutil.Now()
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(solutionCollectionName).Insert(solution); err != nil {
		return err
	}
	return nil
}
