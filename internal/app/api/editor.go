package api

import (
	"github.com/pthethanh/robusta/internal/app/editor"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/upload/cloudinary"
)

func newEditorHandler() (*editor.Handler, error) {
	var conf editor.Config
	envconfig.LoadWithPrefix("EDITOR", &conf)
	cloudinaryConf := cloudinary.LoadConfigFromEnv()
	cloudinaryUploader := cloudinary.New(cloudinaryConf)
	srv := editor.NewService(conf)
	srv.UseUploader(cloudinaryUploader)

	handler := editor.New(conf, srv)
	return handler, nil
}
