package status

import (
	"os"
	"sync"

	"github.com/pthethanh/robusta/internal/pkg/status"

	"gopkg.in/yaml.v2"
)

type (
	Status    = status.Status
	GenStatus struct {
		Success    Status
		NotFound   Status
		Timeout    status.Timeout
		BadRequest Status
		Internal   Status
	}

	EditorStatus struct {
		FileSizeExceedLimit Status `yaml:"file_size_exceed_limit"`
	}

	PolicyStatus struct {
		Unauthorized Status
	}

	ArticleStatus struct {
	}

	UserStatus struct {
		DuplicatedEmail Status
	}

	AuthStatus struct {
		InvalidUserPassword Status `yaml:"invalid_user_password"`
	}

	ChallengeStatus struct {
		NotSupported Status
	}

	statuses struct {
		Gen       GenStatus
		Article   ArticleStatus
		User      UserStatus
		Auth      AuthStatus
		Policy    PolicyStatus
		Editor    EditorStatus
		Challenge ChallengeStatus
	}
)

var (
	all  *statuses
	once sync.Once
)

// Init load statuses from the given config file.
// Init panics if cannot access or error while parsing the config file.
func Init(conf string) {
	once.Do(func() {
		f, err := os.Open(conf)
		if err != nil {
			panic(err)
		}
		all = &statuses{}
		if err := yaml.NewDecoder(f).Decode(all); err != nil {
			panic(err)
		}
	})
}

// all return all registered statuses.
// all will load statuses from configs/Status.yml if the statuses has not initalized yet.
func load() *statuses {
	conf := os.Getenv("STATUS_PATH")
	if conf == "" {
		conf = "configs/status.yml"
	}
	if all == nil {
		Init(conf)
	}
	return all
}

func Gen() GenStatus {
	return load().Gen
}

func Article() ArticleStatus {
	return load().Article
}

func Policy() PolicyStatus {
	return load().Policy
}

func User() UserStatus {
	return load().User
}

func Success() Status {
	return Gen().Success
}

func Editor() EditorStatus {
	return load().Editor
}

func Challenge() ChallengeStatus {
	return load().Challenge
}

func Auth() AuthStatus {
	return load().Auth
}
