package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

const (
	SessionTypeFile = "filesystem"
	SessionTypeMem  = "memory"
	oneWeek         = 7 * 24 * 3600
)

func createSessionStore(conf Config) sessions.Store {
	k := []byte(conf.Session.Secret)
	f := func() sessions.Store {
		s := sessions.NewCookieStore(k)
		s.Options = &sessions.Options{
			HttpOnly: true,
			MaxAge:   oneWeek,
		}
		return s
	}
	switch conf.Session.Type {
	case SessionTypeFile:
		if err := os.MkdirAll(conf.Session.FilePath, os.ModePerm); err != nil {
			log.Errorf("auth config: failed to create sessions file system store, fallback to use cookie store")
			return f()
		}
		s := sessions.NewFilesystemStore(conf.Session.FilePath, k)
		s.Options = &sessions.Options{
			HttpOnly: true,
			MaxAge:   oneWeek,
		}
		return s
	case SessionTypeMem:
		return f()
	default:
		return f()
	}
}
