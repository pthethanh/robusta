package playground

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/uuid"
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

// Run run the code in the target playground server.
func (c *Client) Run(ctx context.Context, r *RunRequest) (*RunResponse, error) {
	code := url.QueryEscape(r.Code)
	playURL := fmt.Sprintf("%s/compile?version=%d&body=%s", c.conf.Host, 2, code)
	req, err := http.NewRequest(http.MethodPost, playURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request to playground server")
	}
	defer res.Body.Close()
	var v RunResponse
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}
	return &v, err
}

// Evaluate evalute the given solution against Go lint rules and run the test.
func (c *Client) Evaluate(ctx context.Context, r *EvaluateRequest) (*EvaluateResponse, error) {
	f, err := MergePackageFiles("main", generatedFileName(), map[string]io.Reader{
		generatedFileName(): bytes.NewBuffer(r.Solution),
		generatedFileName(): bytes.NewBuffer(r.Test),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to merge files")
	}
	problems, err := LintFile(generatedFileName(), f)
	res, err := c.Run(ctx, &RunRequest{
		Code: string(f),
	})
	if err != nil {
		return &EvaluateResponse{
			Problems:     problems,
			IsTestFailed: false,
			Error:        err.Error(),
		}, err
	}
	return &EvaluateResponse{
		Problems:     problems,
		IsTestFailed: res.TestsFailed > 0,
		Error:        res.Errors,
	}, nil
}

// return generated file name.
func generatedFileName() string {
	return fmt.Sprintf("%v.go", uuid.New())
}
