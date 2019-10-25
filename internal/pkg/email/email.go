package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"

	"gopkg.in/gomail.v2"

	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
)

type (
	Sender interface {
		Send(ctx context.Context, email Email) error
	}

	Mailer struct {
		conf   Config
		dialer *gomail.Dialer
	}

	Email struct {
		From    string
		To      []string
		CC      []string
		Subject string
		Body    string
	}

	Config struct {
		Address     string `envconfig:"SMTP_ADDRESS"`
		Username    string `envconfig:"SMTP_USERNAME"`
		Password    string `envconfig:"SMTP_PASSWORD"`
		DefaultFrom string `envconfig:"SMTP_DEFAULT_FROM" default:"gowaynotification@gmail.com"`
	}
)

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func New(conf Config) (*Mailer, error) {
	username := conf.Username
	password := conf.Password
	host, port, _ := net.SplitHostPort(conf.Address)
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("address must be in form of <host>:<port>: %w", err)
	}
	d := gomail.NewDialer(host, portInt, username, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return &Mailer{
		conf:   conf,
		dialer: d,
	}, nil
}

// Send send the give email
func (m *Mailer) Send(ctx context.Context, email Email) error {
	from := email.From
	if from == "" {
		from = m.conf.DefaultFrom
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", email.To...)
	for _, addr := range email.CC {
		msg.SetAddressHeader("Cc", addr, "")
	}
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/html", email.Body)
	if err := m.dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send mail: %w", err)
	}
	return nil
}
