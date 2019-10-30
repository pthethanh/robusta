package status

import (
	"errors"
	"fmt"
)

type (
	Status struct {
		XCode    uint32 `json:"code" yaml:"code"`
		XStatus  int    `json:"status" yaml:"status"`
		XMessage string `json:"message" yaml:"message"`
	}
)

// New return a new status.
func New(code uint32, status int, message string) Status {
	return Status{
		XCode:    code,
		XStatus:  status,
		XMessage: message,
	}
}

func (s Status) Error() string {
	return s.XMessage
}

func (s Status) Code() uint32 {
	return s.XCode
}

func (s Status) Message() string {
	return s.XMessage
}

func (s Status) Status() int {
	return s.XStatus
}

// Is implement errors.Is method
func (s Status) Is(err error) bool {
	var status Status
	if !errors.As(err, &status) {
		return false
	}
	return status.Code() == s.Code()
}

func (s Status) String() string {
	return fmt.Sprintf("code: %d, status: %d, message: %s", s.XCode, s.XStatus, s.XMessage)
}

func (s Status) GoString() string {
	return fmt.Sprintf("code: %d, status: %d, message: %s", s.XCode, s.XStatus, s.XMessage)
}
