package types

import "time"

type (
	User struct {
		ID        string   `json:"id" bson:"_id"`
		Username  string   `json:"username" bson:"username"`
		Password  string   `json:"password" bson:"password"`
		Avatar    string   `json:"avatar" bson:"avatar"`
		FirstName string   `json:"first_name" bson:"first_name"`
		LastName  string   `json:"last_name" bson:"last_name"`
		Status    string   `json:"status" bson:"status"`
		Roles     []string `json:"roles" bson:"roles"`
		Groups    []string `json:"groups" bson:"groups"`

		CreatedAt time.Time `json:"created_at" bson:"created_at"`
		UpdateAt  time.Time `json:"updated_at" bson:"updated_at"`
	}
)
