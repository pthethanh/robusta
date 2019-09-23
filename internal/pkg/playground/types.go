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
		Code   int    `json:"Code"`
		Errors string `json:"Errors"`
		Events []struct {
			Message string `json:"Message"`
			Kind    string `json:"Kind"`
			Delay   int64  `json:"Delay"`
		} `json:"Events"`
		Status      int  `json:"status"`
		IsTest      bool `json:"IsTest"`
		TestsFailed int  `json:"TestsFailed"`
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
		Problems     []lint.Problem `json:"problems"`
		IsTestFailed bool           `json:"is_test_failed"`
		Error        string         `json:"error"`
		TestsFailed  int            `json:"tests_failed"`
	}
)
