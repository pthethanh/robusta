package types

import "encoding/json"

type (
	// BaseResponse is standard response of the app which include code, message, data and meta,...
	BaseResponse struct {
		AppError
		Data interface{} `json:"data"`
		Meta string      `json:"meta"`
	}

	baseResponse BaseResponse

	// IDResponse is a response helper that has ID
	IDResponse struct {
		ID string `json:"id"`
	}
)

var (
	// SuccessResponse is a response with success code
	SuccessResponse = BaseResponse{
		AppError: AppSuccess,
	}
)

// MarshalJSON implement encoding/json.Marshaler interface.
// It will automatically set AppError to Success if AppError is nil
func (rs BaseResponse) MarshalJSON() ([]byte, error) {
	var v = baseResponse(rs)
	if v.AppError.Code() == 0 {
		v.AppError = AppSuccess
	}
	return json.Marshal(v)
}
