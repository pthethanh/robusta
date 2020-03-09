package status_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/pthethanh/robusta/internal/pkg/status"
)

func TestStatusMarshal(t *testing.T) {
	s := status.New(200, 200, "ok")
	b, err := json.Marshal(s)
	if err != nil {
		t.Error(err)
		return
	}
	wants := `{"code":200,"message":"ok","status":200}`
	got := string(b)
	if strings.Contains(got, `"code":200"`) && strings.Contains(got, `"status":200`) {
		t.Errorf("got %s, wants %s", got, wants)
	}
}

func TestStatusUnmarshal(t *testing.T) {
	msg := `{"code":200,"message":"ok","status":200}`
	var s status.Status
	if err := json.Unmarshal([]byte(msg), &s); err != nil {
		t.Error(err)
		return
	}
	wants := status.New(200, 200, "ok")
	if s.Code() != wants.Code() || s.Message() != wants.Message() || s.Status() != wants.Status() {
		t.Errorf("got %v, wants %v", s, wants)
	}
}
