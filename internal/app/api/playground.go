package api

import (
	"github.com/pthethanh/robusta/internal/app/challenge"
	"github.com/pthethanh/robusta/internal/app/playground"
	"github.com/pthethanh/robusta/internal/app/solution"
	client "github.com/pthethanh/robusta/internal/pkg/playground"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func newPlaygroundHandler(challenge playground.ChallengeService, solution playground.SolutionService) *playground.Handler {
	runner := client.New(client.LoadConfigFromEnv())
	srv := playground.NewService(runner, challenge, solution)
	return playground.New(srv)
}

func newChallengeHandler(policy challenge.PolicyService) (*challenge.Handler, *challenge.Service, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, nil, closer, err
	}
	closer.Append(mongoCloser)
	repo := challenge.NewMongoDBRepository(s)
	conf := challenge.LoadConfigFromEnv()
	srv := challenge.NewService(conf, repo, policy)
	handler := challenge.NewHandler(srv)

	return handler, srv, closer, nil
}

func newSolutionHandler(policy solution.PolicyService) (*solution.Handler, *solution.Service, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, nil, closer, err
	}
	closer.Append(mongoCloser)
	repo := solution.NewMongoDBRepository(s)
	conf := solution.LoadConfigFromEnv()
	srv := solution.NewService(conf, repo, policy)
	handler := solution.NewHandler(srv)

	return handler, srv, closer, nil
}
