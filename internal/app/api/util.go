package api

import (
	"os"
	"sync"

	"github.com/pthethanh/robusta/internal/pkg/db/mongodb"
	"github.com/pthethanh/robusta/internal/pkg/util/closeutil"

	"github.com/globalsign/mgo"
)

func staticPath() string {
	pth := os.Getenv("STATIC_PATH")
	if pth != "" {
		return pth
	}
	return "web/dist"
}

func staticPrefix() string {
	prefix := os.Getenv("STATIC_PREFIX")
	if prefix != "" {
		return prefix
	}
	return "/static/"
}

var (
	session     *mgo.Session
	sessionOnce sync.Once
)

func dialDefaultMongoDB() (*mgo.Session, *closeutil.Closer, error) {
	closer := closeutil.NewCloser()
	repoConf := mongodb.LoadConfigFromEnv()
	var err error
	sessionOnce.Do(func() {
		session, err = mongodb.Dial(repoConf)
	})
	if err != nil {
		return nil, closer, err
	}
	s := session.Clone()
	closer.AddFunc(s.Close)
	return s, closer, nil
}
