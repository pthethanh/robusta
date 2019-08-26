package upload

import (
	"io"
	"time"
)

type (
	ResourceType string
)

const (
	Image ResourceType = "image"
	Raw   ResourceType = "raw"
	Video ResourceType = "video"
	Auto  ResourceType = "auto"
)

type (
	ByURLRequest struct {
		Type     ResourceType
		PublicID string
		URL      string
	}

	Request struct {
		Type     ResourceType
		PublicID string
		Reader   io.Reader
	}

	Response struct {
		Bytes            int64     `json:"bytes"`
		CreatedAt        time.Time `json:"created_at"`
		Etag             string    `json:"etag"`
		Format           string    `json:"format"`
		Height           int       `json:"height"`
		OriginalFilename string    `json:"original_filename"`
		Placeholder      bool      `json:"placeholder"`
		PublicID         string    `json:"public_id"`
		ResourceType     string    `json:"resource_type"`
		SecureURL        string    `json:"secure_url"`
		Signature        string    `json:"signature"`
		Tags             []string  `json:"tags"`
		Type             string    `json:"type"`
		URL              string    `json:"url"`
		Version          int64     `json:"version"`
		Width            int       `json:"width"`
	}
)
