package htmlutil

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/pkg/log"
)

// LoadTemplates load a list of templates in the target folder.
// Sub folders is not supported.
func LoadTemplates(folder string) (*template.Template, error) {
	f, err := os.Open(folder)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open target folder")
	}
	infos, err := f.Readdir(-1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read target folder")
	}
	var templates []string
	for _, info := range infos {
		p := filepath.Join(folder, info.Name())
		if info.IsDir() {
			log.Warnf("ignore sub folder: %s", p)
			continue
		}
		templates = append(templates, p)
	}
	t, err := template.ParseFiles(templates...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse templates")
	}
	return t, nil
}
