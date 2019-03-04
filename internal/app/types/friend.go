package types

// Friend hold information of a friend
type Friend struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name"`
}
