package policy

type (
	AssignPolicyRequest struct {
		Subject string `json:"subject" validate:"required"`
		Object  string `json:"object" validate:"required"`
		Action  string `json:"action" validate:"required"`
		Effect  string `json:"effect" validate:"oneof=* allow deny"`
	}

	AssignGroupPolicyRequest struct {
		Subject string `json:"subject" validate:"required"`
		Group   string `json:"group" validate:"group"`
	}

	CreateGroupRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)

const (
	ActionPolicyUpdate = "policy:update"
	ActionPolicyRead   = "policy:read"
	Object             = "policy"
)
