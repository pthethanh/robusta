package types

import (
	"errors"
	"net/http"
)

// AppError supports for general cases
var (
	AppCodeSuccess = http.StatusOK

	AppSuccess = AppError{
		XCode:    http.StatusOK,
		XMessage: "success",
	}
	ErrNotFound     = NewAppError(http.StatusNotFound, errors.New("notfound"))
	ErrUnauthorized = NewAppError(http.StatusUnauthorized, errors.New("unauthorized"))
)

// NewAppError return new AppError instance which implements error interface
func NewAppError(code int, err error) AppError {
	msg := "success"
	if err != nil {
		msg = err.Error()
	}
	return AppError{
		err:      err,
		XCode:    code,
		XMessage: msg,
	}
}

// AppError implement error and AppMessage interface to provide additional information in error report.
type AppError struct {
	err      error
	XCode    int    `json:"code"`
	XMessage string `json:"message"`
}

// Error implement error interface
func (err AppError) Error() string {
	if err.err == nil {
		return ""
	}
	return err.err.Error()
}

// Code implement AppMessage interface
func (err AppError) Code() int {
	return err.XCode
}

// Message implement AppMessage interface
func (err AppError) Message() string {
	return err.XMessage
}
