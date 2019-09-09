package playground

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
)

type (
	Client struct {
		client *http.Client
		conf   Config
	}

	Config struct {
		Timeout time.Duration `envconfig:"PLAYGROUND_TIMEOUT" default:"30s"`
		Host    string        `envconfig:"PLAYGROUND_HOST" default:"https://play.golang.org"`
	}
)

func New(conf Config) *Client {
	return &Client{
		client: &http.Client{
			Timeout: conf.Timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		conf: conf,
	}
}

// LoadConfigFromEnv load configuration from environment variables.
func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

// Run run the code in the target playground server
func (s *Client) Run(ctx context.Context, r *Request) (*Response, error) {
	code := url.QueryEscape(r.Code)
	playURL := fmt.Sprintf("%s/compile?version=%d&body=%s", s.conf.Host, 2, code)
	req, err := http.NewRequest(http.MethodPost, playURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	res, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request to playground server")
	}
	defer res.Body.Close()
	var v Response
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}
	return &v, err
}
