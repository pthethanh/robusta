package api

import "github.com/pthethanh/robusta/internal/pkg/email"

func createMailer() *email.Mailer {
	conf := email.LoadConfigFromEnv()
	mailer, err := email.New(conf)
	if err != nil {
		panic(err)
	}
	return mailer
}
