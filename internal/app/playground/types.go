package playground

import "github.com/pthethanh/robusta/internal/pkg/playground"

type (
	Request struct {
		Code string `json:"code"`
	}

	Response = playground.RunResponse

	EvaluateRequest struct {
		ChallengeID string `json:"challenge_id" validate:"required"`
		FolderID    string `json:"folder_id" validate:"required"`
		Solution    string `json:"solution" validate:"min=15"`
	}
)
