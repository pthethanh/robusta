package api

import (
	"github.com/pthethanh/robusta/internal/app/folder"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"
)

func newFolderHandler(policy folder.PolicyService) (*folder.Handler, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	s, mongoCloser, err := dialDefaultMongoDB()
	if err != nil {
		return nil, closer, err
	}
	closer.Append(mongoCloser)
	repo := folder.NewMongoDBRepository(s)
	conf := folder.LoadConfigFromEnv()
	srv := folder.NewService(conf, repo, policy)
	handler := folder.NewHandler(srv)

	return handler, closer, nil
}
