package mock

import (
	"net/http"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/respond"
	"github.com/pthethanh/robusta/internal/pkg/uuid"
)

type (
	article struct {
		ID            string    `json:"id"`
		Title         string    `json:"title"`
		Content       string    `json:"content"`
		Abstract      string    `json:"abstract"`
		CreatedByID   string    `json:"created_by_id"`
		CreatedByName string    `json:"created_by_name"`
		Views         int64     `json:"views"`
		Likes         int64     `json:"likes"`
		Comments      int64     `json:"comments"`
		CreatedAt     time.Time `json:"created_at"`
	}
)

func ArticleList(w http.ResponseWriter, r *http.Request) {
	n := 50
	result := make([]article, 0)
	for i := 0; i < n; i++ {
		result = append(result, article{
			ID:    uuid.New(),
			Title: "Go " + time.Now().String(),
			Abstract: `<p>
			A sense of community flourishes when we come together in person. As handles become names and avatars become faces, the smiles are real and true friendship can grow. There is joy in the sharing of knowledge and celebrating the accomplishments of our friends, colleagues, and neighbors. In our rapidly growing Go community this critical role is played by the Go user groups.
			
			To better support our Go user groups worldwide, the Go community leaders at GoBridge and Google have joined forces to create a new program called the Go Developer Network (GDN). The GDN is a collection of Go user groups working together with a shared mission to empower developer communities with the knowledge, experience, and wisdom to build the next generation of software in Go.</p>`,
			Content:       "Go content",
			CreatedByID:   uuid.New()[:10],
			CreatedByName: uuid.New()[:15],
			Views:         100,
			Comments:      200,
			Likes:         500,
			CreatedAt:     time.Now(),
		})
	}
	respond.JSON(w, http.StatusOK, map[string]interface{}{
		"code":    20000,
		"message": "success",
		"items":   result,
	})
}
