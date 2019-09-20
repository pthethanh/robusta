package challenge

type (
	// FindRequest hold information of finding challenges.
	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		Tags        []string `json:"tags" validate:"required_without=CreatedByID IDs"`
		CreatedByID string   `json:"created_by_id" validate:"required_without=IDs Tags"`
		IDs         []string `json:"ids" validate:"required_without=CreatedByID Tags"`
		SortBy      []string `json:"sort_by"`
		FolderID    string   `json:"folder_id" validate:"required"`
	}
)

// Actions policy
const (
	ActionCreate = "challenge:create"
	ActionDelete = "challenge:delete"
	ActionUpdate = "challenge:update"
	ActionRead   = "challenge:read"
)

// Policy object name
const (
	PolicyObject = "challenge"
)
