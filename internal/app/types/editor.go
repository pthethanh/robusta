package types

type (
	EditorJSBlock struct {
		Type string                 `json:"type" bson:"type" validate:"required"`
		Data map[string]interface{} `json:"data" bson:"data"`
	}
	EditorJSContent struct {
		Time    int64           `json:"time,omitempty" bson:"time" validate:"required"`
		Blocks  []EditorJSBlock `json:"blocks,omitempty" bson:"blocks" validate:"required,gte=1"`
		Version string          `json:"version,omitempty" bson:"version" validate:"required"`
	}
)
