package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/htmlutil"
)

type (
	IndexHandler struct {
	}
)

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	indexHTML := fmt.Sprintf("%s/index.html", staticPath())
	f, err := os.Open(indexHTML)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	defer f.Close()
	if ok, _ := strconv.ParseBool(os.Getenv("HTTP_ENABLE_PUSH")); !ok {
		// directly
		if _, err := io.Copy(w, f); err != nil {
			respond.Error(w, err, http.StatusInternalServerError)
			return
		}
	}
	buff := &bytes.Buffer{}
	if _, err := io.Copy(buff, f); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	// update index and static resource list
	resources, err := htmlutil.ReadHeaderLinks(buff)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to read header links, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
	}
	if pusher, ok := w.(http.Pusher); ok {
		for _, rs := range resources {
			if err := pusher.Push(rs.Href, nil); err != nil {
				log.WithContext(r.Context()).Errorf("failed to push, rs: %s, err: %v", rs, err)
			} else {
				log.WithContext(r.Context()).Infof("pushed: %s", rs)
			}
		}
	}
	// return to client
	if _, err := io.Copy(w, buff); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
}
