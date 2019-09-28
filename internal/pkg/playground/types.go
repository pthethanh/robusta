package playground

import (
	"context"

	"golang.org/x/lint"
)

type (
	RunRequest struct {
		Code string `json:"code"`
	}

	RunResponse struct {
		Code        int     `json:"code"`
		Errors      string  `json:"errors"`
		Events      []Event `json:"events"`
		Status      int     `json:"status"`
		IsTest      bool    `json:"is_test"`
		TestsFailed int     `json:"tests_failed"`
	}

	Event struct {
		Message string `json:"message"`
		Kind    string `json:"kind"`
		Delay   int64  `json:"delay"`
	}

	Runner interface {
		// Run run the code in the target playground server.
		Run(ctx context.Context, r *RunRequest) (*RunResponse, error)
		// Evaluate evalute the given solution against Go lint rules and run the test.
		Evaluate(ctx context.Context, r *EvaluateRequest) (*EvaluateResponse, error)
	}

	EvaluateRequest struct {
		Solution []byte
		Test     []byte
	}

	EvaluateResponse struct {
		Status       int            `json:"status"`
		Events       []Event        `json:"events"`
		Problems     []lint.Problem `json:"problems"`
		IsTestFailed bool           `json:"is_test_failed"`
		Error        string         `json:"error"`
		TestsFailed  int            `json:"tests_failed"`
	}
)

// IsSuccess report whether the evaluate result is success or failed.
func (rs EvaluateResponse) IsSuccess() bool {
	return rs.Status == 0 && rs.Error == "" && rs.TestsFailed == 0
}
