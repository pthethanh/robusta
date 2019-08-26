package types

import "time"

type (
	TutorialStep struct {
		Title       string      `json:"title" bson:"title"`
		Content     Content     `json:"content" bson:"content"`
		ContentType ContentType `json:"content_type" bson:"content_type"`
	}
	Tutorial struct {
		ID            string         `json:"id" bson:"_id"`
		Title         string         `json:"title" bson:"title"`
		Steps         []TutorialStep `json:"steps"`
		Source        string         `json:"source" bson:"source"`
		CreatedByID   string         `json:"created_by_id" bson:"created_by_id"`
		CreatedByName string         `json:"created_by_name" bson:"created_by_name"`
		Status        Status         `json:"status" bson:"status"`
		PublishDate   *time.Time     `json:"publish_date" bson:"publish_date"`
		Promoted      bool           `json:"promoted" bson:"promoted"`
		PromotedRank  int            `json:"promoted_rank" bson:"promoted_rank"`
		Views         int64          `json:"views" bson:"views"`
		Likes         int64          `json:"likes" bson:"likes"`
		Comments      int64          `json:"comments" bson:"comments"`
		CreatedAt     *time.Time     `json:"created_at" bson:"created_at"`
		UpdatedAt     *time.Time     `json:"update_at" bson:"updated_at"`
	}
)
