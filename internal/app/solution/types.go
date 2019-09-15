package solution

type (
	FindRequest struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		ChallengeID string   `json:"challenge_id"`
		CreatedByID string   `json:"created_by_id"`
		SortBy      []string `json:"sort_by"`
	}
)
