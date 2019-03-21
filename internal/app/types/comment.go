package types

import "time"

type (
	Comment struct {
		ID              string `json:"id,omitempty" bson:"_id"`
		Target          string `json:"target,omitempty" bson:"target"`
		Content         string `json:"content,omitempty" bson:"content"`
		CreatedByName   string `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByID     string `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedByAvatar string `json:"created_by_avatar" bson:"created_by_avatar"`
		Parent          string `json:"parent,omitempty" bson:"parent"`

		CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt time.Time `json:"modified_at,omitempty" bson:"updated_at"`
	}
)
