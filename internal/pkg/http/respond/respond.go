package respond

import (
	"encoding/json"
	"net/http"
)

type (
	appError interface {
		Error() string
		Code() uint32
		Message() string
	}
)

// JSON write status and JSON data to http response writer
func JSON(w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// Error write error, status to http response writer with JSON format: {"code": status, "message": error}
func Error(w http.ResponseWriter, err error, status int) {
	if appError, ok := err.(appError); ok {
		JSON(w, status, appError)
		return
	}
	JSON(w, status, map[string]interface{}{
		"code":    status,
		"message": err.Error(),
	})
}
