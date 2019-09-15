package types

import (
	"time"
)

// Challenge types
const (
	ChallengeTypeExercise ChallengeType = "exercise"
)

// Programming languages
const (
	ProgrammingLanguageGo ProgrammingLanguage = "go"
)

type (
	ChallengeType       string
	ProgrammingLanguage string
	Challenge           struct {
		ID                  string              `json:"id" bson:"_id"`
		Title               string              `json:"title" bson:"title"`
		Description         string              `json:"description" bson:"description"`
		Type                ChallengeType       `json:"type" bson:"type"`
		ProgrammingLanguage ProgrammingLanguage `json:"programming_language" bson:"programming_language"`
		Tags                []string            `json:"tags" bson:"tags"`
		Test                string              `json:"test" bson:"test"`
		CreatedAt           *time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt           *time.Time          `json:"updated_at" bson:"updated_at"`
		CreatedByID         string              `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedByName       string              `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByAvatar     string              `json:"created_by_avatar,omitempty" bson:"created_by_avatar"`
	}
)
