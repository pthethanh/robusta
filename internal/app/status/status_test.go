package status_test

import (
	"testing"

	"github.com/pthethanh/robusta/internal/app/status"
)

func TestInit(t *testing.T) {
	status.Init("../../../configs/status.yml")
	if status.Success().Code() != 1000 {
		t.Errorf("got success=%d, wants %d", status.Success().Code(), 1000)
	}
}
