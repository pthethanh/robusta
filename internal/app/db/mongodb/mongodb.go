package mongodb

import (
	"time"

	"github.com/pthethanh/robusta/internal/pkg/glog"

	"github.com/globalsign/mgo"
)

type (
	// Config hold MongoDB configuration information
	Config struct {
		Addrs    []string      `envconfig:"MONGODB_ADDRS" default:"127.0.0.1:27017"`
		Database string        `envconfig:"MONGODB_DATABASE" default:"robusta"`
		Username string        `envconfig:"MONGODB_USERNAME"`
		Password string        `envconfig:"MONGODB_PASSWORD"`
		Timeout  time.Duration `envconfig:"MONGODB_TIMEOUT" default:"10s"`
	}
)

// Dial dial to target server with Monotonic mode
func Dial(conf *Config, logger glog.Logger) (*mgo.Session, error) {
	logger.Infof("dialing to target MongoDB at: %v, database: %v", conf.Addrs, conf.Database)
	ms, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    conf.Addrs,
		Database: conf.Database,
		Username: conf.Username,
		Password: conf.Password,
		Timeout:  conf.Timeout,
	})
	if err != nil {
		return nil, err
	}
	ms.SetMode(mgo.Monotonic, true)
	logger.Infof("successfully dialing to MongoDB at %v", conf.Addrs)
	return ms, nil
}
