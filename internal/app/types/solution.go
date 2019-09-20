package types

import (
	"time"
)

// Solution statuses
const (
	SolutionStatusSuccess SolutionStatus = "success"
	SolutionStatusFailed  SolutionStatus = "failed"
)

type (
	SolutionStatus string
	Solution       struct {
		ID              string         `json:"id" bson:"_id"`
		ChallengeID     string         `json:"challenge_id" bson:"challenge_id"`
		Content         string         `json:"content" bson:"content" validate:"required"`
		Status          SolutionStatus `json:"status" bson:"status" validate:"required"`
		EvaluateResult  string         `json:"evaluate_result" bson:"evaluate_result"`
		CreatedAt       *time.Time     `json:"created_at" bson:"created_at"`
		CreatedByID     string         `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedByName   string         `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByAvatar string         `json:"created_by_avatar,omitempty" bson:"created_by_avatar"`
	}
)
