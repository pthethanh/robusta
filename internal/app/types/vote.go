package types

import "time"

type (
	Vote struct {
		ID            string     `json:"id,omitempty" bson:"_id"`
		Target        string     `json:"target,omitempty" bson:"target"`
		CreatedByName string     `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByID   string     `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedAt     *time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	}
)
