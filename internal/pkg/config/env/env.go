package envconfig

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var envPrefix = ""

// Load loads the environment variables into the provided struct
func Load(t interface{}) {
	if err := envconfig.Process(envPrefix, t); err != nil {
		log.Fatalf("config: unable to load config for %T: %s", t, err)
	}
}

//LoadWithPrefix loads the environment variables with prefix into the provided struct
func LoadWithPrefix(prefix string, t interface{}) {
	if err := envconfig.Process(prefix, t); err != nil {
		log.Fatalf("config: unable to load config for %T: %s", t, err)
	}
}
