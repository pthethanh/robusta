package challenge

type (
	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		Tags        []string `json:"tags"`
		CreatedByID string   `json:"created_by_id"`
		SortBy      []string `json:"sort_by"`
	}
)
