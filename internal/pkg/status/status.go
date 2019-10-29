package status

import (
	"encoding/json"
	"net/http"
)

type (
	Status struct {
		code    uint32
		status  int
		message string
	}
)

var (
	Success      = New(CodeSuccess, http.StatusOK, "success")
	Unauthorized = New(CodeUnauthorized, http.StatusUnauthorized, "You're not authorized to access the resource.")
)

// New return a new status.
func New(code uint32, status int, message string) Status {
	return Status{
		code:    code,
		status:  status,
		message: message,
	}
}

func (s Status) Error() string {
	return s.message
}

func (s Status) Code() uint32 {
	return s.code
}

func (s Status) Message() string {
	return s.message
}

func (s Status) Status() int {
	return s.status
}

func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code":    s.code,
		"message": s.message,
		"status":  s.status,
	})
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var m struct {
		Code    uint32
		Status  int
		Message string
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	s.code = m.Code
	s.message = m.Message
	s.status = m.Status
	return nil
}
