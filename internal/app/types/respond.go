package types

import (
	"encoding/json"

	"github.com/pthethanh/robusta/internal/app/status"
)

type (
	// BaseResponse is standard response of the app which include code, message, data and meta,...
	BaseResponse struct {
		status.Status
		Data interface{} `json:"data"`
		Meta string      `json:"meta"`
	}

	baseResponse BaseResponse

	// IDResponse is a response helper that has ID
	IDResponse struct {
		ID string `json:"id"`
	}
)

// MarshalJSON implement encoding/json.Marshaler interface.
// It will automatically set AppError to Success if AppError is nil
func (rs BaseResponse) MarshalJSON() ([]byte, error) {
	var v = baseResponse(rs)
	if v.Status.Status() == 0 {
		v.Status = status.Gen().Success
	}
	return json.Marshal(v)
}
