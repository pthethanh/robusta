package challenge

type (
	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		Tags        []string `json:"tags" validate:"required_without=CreatedByID IDs"`
		CreatedByID string   `json:"created_by_id" validate:"required_without=IDs Tags"`
		IDs         []string `json:"ids" validate:"required_without=CreatedByID Tags"`
		SortBy      []string `json:"sort_by"`
	}
)
