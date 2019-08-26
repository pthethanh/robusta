package handlerutil

import (
	"net/url"
	"strconv"
)

func IntFromQuery(k string, req url.Values, def int) int {
	v := req.Get(k)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}
