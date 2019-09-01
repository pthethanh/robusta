package types

type (
	ArticleInfo struct {
		ID              string `json:"id,omitempty"`
		Title           string `json:"title,omitempty"`
		CreatedByID     string `json:"created_by_id,omitempty"`
		CreatedByName   string `json:"created_by_name,omitempty"`
		CreatedByAvatar string `json:"created_by_avatar,omitempty"`
	}
)
