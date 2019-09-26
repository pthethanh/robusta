package playground

import "github.com/pthethanh/robusta/internal/pkg/playground"

type (
	Request struct {
		Code string `json:"code"`
	}

	Response struct {
		Code        int                `json:"code"`
		Errors      string             `json:"errors"`
		Events      []playground.Event `json:"events"`
		Status      int                `json:"status"`
		IsTest      bool               `json:"is_test"`
		TestsFailed int                `json:"tests_failed"`
	}

	EvaluateRequest struct {
		ChallengeID string `json:"challenge_id" validate:"required"`
		FolderID    string `json:"folder_id" validate:"required"`
		Solution    string `json:"solution" validate:"min=15"`
	}
)
