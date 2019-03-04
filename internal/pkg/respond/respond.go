package respond

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// JSON write status and JSON data to http response writer
func JSON(w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		Error(w, errors.Wrap(err, "json marshal failed"), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(b)
}

// Error write error, status to http response writer
func Error(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}
