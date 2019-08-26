package editor

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/image"
	"github.com/pthethanh/robusta/internal/pkg/linkresolver"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/upload"
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
			AppError: types.AppSuccess,
			Success:  LinkStatusSuccess,
			Meta: LinkMeta{
				Title: url,
			},
		}, nil
	}
	return &Link{
		AppError: types.AppSuccess,
		Success:  LinkStatusSuccess,
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
		return "", errors.Wrap(err, "failed to download image")
	}
	defer res.Body.Close()
	r := io.LimitReader(res.Body, s.conf.MaxUploadSize)
	name := strings.Split(path.Base(url), ".")[0] // just file name
	return s.uploadImages(ctx, name, r)
}

func (s *Service) UploadImageByFile(ctx context.Context, name string, r io.Reader) (string, error) {
	nameOnly := strings.Split(name, ".")[0]
	return s.uploadImages(ctx, nameOnly, r)
}

func (s *Service) uploadImages(ctx context.Context, name string, r io.Reader) (string, error) {
	web := &bytes.Buffer{}
	// resize the image
	if _, err := s.resizeImage(r, image.ResizeTarget{
		Option: image.ResizeOptionWeb,
		Writer: web,
	}); err != nil {
		return "", errors.Wrap(err, "failed to resize images")
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
