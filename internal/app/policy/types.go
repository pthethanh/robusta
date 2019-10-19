package policy

type (
	GroupPolicy struct {
		Subject string `json:"subject" validate:"required"`
		Group   string `json:"group" validate:"group"`
	}

	FindPolicyRequest struct {
		Subjects []string `json:"subject"`
		Actions  []string `json:"actions"`
		Objects  []string `json:"objects"`
	}
)

const (
	ActionPolicyUpdate = "policy:update"
	ActionPolicyRead   = "policy:read"
	Object             = "policy"
)
