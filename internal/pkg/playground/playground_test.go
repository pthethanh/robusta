// +build integration

package playground_test

import (
	"context"
	"testing"

	"github.com/pthethanh/robusta/internal/pkg/playground"
)

func TestEvaluate(t *testing.T) {
	solutionFile := `
	package test

	func Sum(a, b int) int {
		return a + b
	}`

	testFile := `
	package test

	import (
		"testing"
	)
	
	func TestSum(t *testing.T) {
		s := Sum(1, 2)
		if s != 3 {
			t.Fail()
		}
	}`

	conf := playground.LoadConfigFromEnv()
	c := playground.New(conf)
	res, err := c.Evaluate(context.Background(), &playground.EvaluateRequest{
		Solution: []byte(solutionFile),
		Test:     []byte(testFile),
	})
	if err != nil {
		t.Error(err)
		return
	}
	if res.IsTestFailed {
		t.Errorf("got Evaluate result failed, want success; err: %s", res.Error)
	}
}
