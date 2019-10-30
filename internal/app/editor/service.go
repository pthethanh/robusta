package editor

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/pkg/image"
	"github.com/pthethanh/robusta/internal/pkg/linkresolver"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/upload"
	"github.com/pthethanh/robusta/internal/pkg/uuid"
)

type (
	Uploader interface {
		UploadByURL(ctx context.Context, req upload.ByURLRequest) (*upload.Response, error)
		Upload(ctx context.Context, req upload.Request) (*upload.Response, error)
	}

	ImageResizeFunc = func(r io.Reader, targets ...image.ResizeTarget) (string, error)

	Service struct {
		linkResolver *linkresolver.Resolver
		uploader     Uploader
		resizeImage  ImageResizeFunc
		conf         Config
		httpClient   *http.Client
	}
)

func NewService(conf Config) *Service {
	client := &http.Client{
		Timeout: conf.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &Service{
		linkResolver: linkresolver.New(),
		conf:         conf,
		resizeImage:  image.Resize,
		httpClient:   client,
	}
}

func (s *Service) UseUploader(uploader Uploader) {
	s.uploader = uploader
}

// FetchURL fetch information like title, image, description of the link
// Even it's failed, we want to return the original link
// without description or image
func (s *Service) FetchURL(ctx context.Context, url string) (*Link, error) {
	link, err := s.linkResolver.Resolve(url)
	if err != nil {
		return &Link{
			Status:  status.Success(),
			Success: LinkStatusSuccess,
			Meta: LinkMeta{
				Title: url,
			},
		}, nil
	}
	return &Link{
		Status:  status.Success(),
		Success: LinkStatusSuccess,
		Meta: LinkMeta{
			Title:       link.Title,
			Description: link.Description,
			Image: LinkImage{
				URL: link.Image.URL,
			},
		},
	}, nil
}

func (s *Service) UploadImageByURL(ctx context.Context, url string) (string, error) {
	res, err := s.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %w", err)
	}
	defer res.Body.Close()
	r := io.LimitReader(res.Body, s.conf.MaxUploadSize)
	name := uuid.New()
	return s.uploadImages(ctx, name, r)
}

func (s *Service) UploadImageByFile(ctx context.Context, name string, r io.Reader) (string, error) {
	newName := uuid.New()
	return s.uploadImages(ctx, newName, r)
}

func (s *Service) uploadImages(ctx context.Context, name string, r io.Reader) (string, error) {
	web := &bytes.Buffer{}
	// resize the image
	if _, err := s.resizeImage(r, image.ResizeTarget{
		Option: image.ResizeOptionWeb,
		Writer: web,
	}); err != nil {
		return "", fmt.Errorf("failed to resize images: %w", err)
	}
	// upload the image to remote storage
	res, err := s.uploader.Upload(ctx, upload.Request{
		Type:     upload.Image,
		PublicID: name,
		Reader:   web,
	})
	if err != nil {
		log.WithContext(ctx).Errorf("failed to upload image from file, url: %s, err: %v", name, err)
		return "", err
	}
	log.WithContext(ctx).Debugf("upload res: %v", res)
	return res.URL, nil
}
