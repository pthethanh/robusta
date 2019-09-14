package playground

import (
	"context"
	"golang.org/x/lint"
)

type (
	Request struct {
		Code string `json:"code"`
	}

	Response struct {
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
		Run(ctx context.Context, r *Request) (*Response, error)
	}

	EvaluateRequest struct {
		Solution []byte
		Test     []byte
	}

	EvaluateResponse struct {
		Problems     []lint.Problem
		IsTestFailed bool
		Error        string
	}
)
