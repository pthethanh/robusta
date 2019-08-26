package playground

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

const (
	playGroundURL = "https://play.golang.org/compile"
)

type (
	Service struct {
		client     *http.Client
		serviceURL string
	}
)

func NewService() *Service {
	return &Service{
		client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		serviceURL: playGroundURL,
	}
}

func (s *Service) Run(ctx context.Context, r *types.PlaygroundRequest) (*types.PlaygroundResponse, error) {
	code := url.QueryEscape(r.Code)
	playURL := fmt.Sprintf("%s?version=%d&body=%s", s.serviceURL, 2, code)
	req, err := http.NewRequest(http.MethodPost, playURL, nil)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to create request, err: %v", err)
		return nil, err
	}
	res, err := s.client.Do(req)
	if err != nil {
		log.WithContext(ctx).Errorf("failed to request to playground, err: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	var v types.PlaygroundResponse
	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return nil, err
	}
	return &v, err
}
