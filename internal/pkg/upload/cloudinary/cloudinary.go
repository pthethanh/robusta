package cloudinary

import (
	"bytes"
	"context"
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/upload"
)

type (
	Config struct {
		Host      string        `envconfig:"CLOUDINARY_HOST" default:"https://api.cloudinary.com/v1_1"`
		APIKey    string        `envconfig:"CLOUDINARY_API_KEY"`
		Secret    string        `envconfig:"CLOUDINARY_SECRET"`
		CloudName string        `envconfig:"CLOUDINARY_CLOUD_NAME"`
		Timeout   time.Duration `envconfig:"CLOUDINARY_TIME_OUT" default:"30s"`
	}
	Client struct {
		httpClient *http.Client
		conf       Config
	}
)

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func New(conf Config) *Client {
	client := &http.Client{
		Timeout: conf.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &Client{
		conf:       conf,
		httpClient: client,
	}
}

// UploadByURL upload resource by its URL
func (c *Client) UploadByURL(ctx context.Context, req upload.ByURLRequest) (*upload.Response, error) {
	var res upload.Response
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	params := fmt.Sprintf("public_id=%s&timestamp=%s%s", req.PublicID, timestamp, c.conf.Secret)
	h := sha1.New()
	h.Write([]byte(params))
	signature := h.Sum(nil)
	r, err := newUploadRequest(c.getUploadURL(req.Type), map[string]string{
		"api_key":   c.conf.APIKey,
		"timestamp": timestamp,
		"signature": hex.EncodeToString(signature),
		"public_id": req.PublicID,
		"file":      req.URL,
	})
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to upload, status: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Upload upload resource base on the given reader: file, multipart reader,...
func (c *Client) Upload(ctx context.Context, req upload.Request) (*upload.Response, error) {
	var res upload.Response
	timestamp := fmt.Sprintf("%d", time.Now().Unix())
	params := fmt.Sprintf("public_id=%s&timestamp=%s%s", req.PublicID, timestamp, c.conf.Secret)
	h := sha1.New()
	h.Write([]byte(params))
	signature := h.Sum(nil)
	r, err := newFileUploadRequest(c.getUploadURL(req.Type), map[string]string{
		"api_key":   c.conf.APIKey,
		"timestamp": timestamp,
		"signature": hex.EncodeToString(signature),
		"public_id": req.PublicID,
	}, "file", req.PublicID, req.Reader)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to upload, status: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *Client) getUploadURL(resourceType upload.ResourceType) string {
	return fmt.Sprintf("%s/%s/%s/upload", c.conf.Host, c.conf.CloudName, resourceType)
}

// Creates a new file upload http request with optional extra params
func newFileUploadRequest(uri string, params map[string]string, paramName string, fileName string, reader io.Reader) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fileName)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, reader); err != nil {
		return nil, err
	}
	for key, val := range params {
		if err := writer.WriteField(key, val); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

// Creates a new file upload http request with optional extra params
func newUploadRequest(uri string, params map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		if err := writer.WriteField(key, val); err != nil {
			return nil, err
		}
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
