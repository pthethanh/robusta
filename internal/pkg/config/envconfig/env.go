package envconfig

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"

	"github.com/pthethanh/robusta/internal/pkg/log"
)

var envPrefix = ""

// Load loads the environment variables into the provided struct
func Load(t interface{}) {
	if err := envconfig.Process(envPrefix, t); err != nil {
		log.Errorf("config: unable to load config for %T: %s", t, err)
	}
}

//LoadWithPrefix loads the environment variables with prefix into the provided struct
func LoadWithPrefix(prefix string, t interface{}) {
	if err := envconfig.Process(prefix, t); err != nil {
		log.Errorf("config: unable to load config for %T: %s", t, err)
	}
}

// SetEnvFromFile load environments from file and set it to system environment via os.Setenv
func SetEnvFromFile(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "#") || strings.TrimSpace(txt) == "" {
			continue
		}
		env := strings.SplitN(txt, "=", 2)
		if len(env) != 2 {
			return errors.New("environment must be in key=value pair")
		}
		k := env[0]
		v := env[1]
		os.Setenv(k, v)
		log.Infof("set env, key=%s", k)
	}
	return nil
}
