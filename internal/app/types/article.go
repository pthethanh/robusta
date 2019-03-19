package types

import (
	"fmt"
	"time"
)

type (
	Article struct {
		ID            string    `json:"id" bson:"_id"`
		Title         string    `json:"title" bson:"title"`
		Content       string    `json:"content" bson:"content"`
		Abstract      string    `json:"abstract" bson:"abstract"`
		CreatedByID   string    `json:"created_by_id" bson:"created_by_id"`
		CreatedByName string    `json:"created_by_name" bson:"created_by_name"`
		Views         int64     `json:"views" bson:"views"`
		Likes         int64     `json:"likes" bson:"likes"`
		Comments      int64     `json:"comments" bson:"comments"`
		CreatedAt     time.Time `json:"created_at" bson:"created_at"`
		UpdatedAt     time.Time `json:"update_at" bson:"updated_at"`
	}
)

func (a *Article) String() string {
	return fmt.Sprintf("Title: %s\nAbstract: %s\nContent: %s\nCreatedByName: %s\nCreatedByID: %s\nCreatedAt: %v\n", a.Title, a.Abstract, a.Content, a.CreatedByName, a.CreatedByID, a.CreatedAt)
}
