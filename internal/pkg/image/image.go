package image

import (
	"fmt"
	"image"
	"io"

	"github.com/disintegration/imaging"
)

type (
	ResizeOption struct {
		MaxWidth  int
		MaxHeight int
	}

	ResizeTarget struct {
		Writer io.Writer
		Option ResizeOption
	}
)

var (
	ResizeOptionThumbnail ResizeOption = ResizeOption{
		MaxWidth: 320,
	}

	ResizeOptionWeb ResizeOption = ResizeOption{
		MaxWidth: 640,
	}
)

// Thumbnail resize the image as thumbnail, see size definition in image.ResizeOptionThumbnail
func Thumbnail(w io.Writer, r io.Reader) (string, error) {
	return Resize(r, ResizeTarget{
		Option: ResizeOptionThumbnail,
		Writer: w,
	})
}

// Web resize the image for web, see size defintion in image.ResizeOptionWeb
func Web(w io.Writer, r io.Reader) (string, error) {
	return Resize(r, ResizeTarget{
		Option: ResizeOptionWeb,
		Writer: w,
	})
}

// Resize resize the given image into multiple targets
func Resize(r io.Reader, targets ...ResizeTarget) (string, error) {
	img, ext, err := image.Decode(r)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}
	size := img.Bounds().Size()
	for _, target := range targets {
		width := size.X
		height := size.Y
		if target.Option.MaxWidth != 0 && width > target.Option.MaxWidth {
			width = target.Option.MaxWidth
			height = 0 // will be calculated base on image ratio
		} else if target.Option.MaxHeight != 0 && height > target.Option.MaxHeight {
			height = target.Option.MaxHeight
			width = 0 // will be calculated base on image ratio
		}
		newImg := imaging.Resize(img, width, height, imaging.Box)
		format, err := imaging.FormatFromExtension(ext)
		if err != nil {
			return "", fmt.Errorf("failed to get file format: %w", err)
		}
		if err := imaging.Encode(target.Writer, newImg, format); err != nil {
			return "", fmt.Errorf("failed to encode image: %w", err)
		}
	}
	return ext, nil
}
