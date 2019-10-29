package types

import (
	"time"
)

// Challenge types
const (
	ChallengeTypeGoExercise ChallengeType = "go_exercise"
)

// Programming languages
const (
	ProgrammingLanguageGo ProgrammingLanguage = "go"
)

const (
	ChallengeLevelEasy ChallengeLevel = 1
)

type (
	ChallengeType       string
	ChallengeLevel      int
	ProgrammingLanguage string
	Challenge           struct {
		ID              string          `json:"id" bson:"_id"`
		Title           string          `json:"title" bson:"title"`
		Description     EditorJSContent `json:"description" bson:"description"`
		Type            ChallengeType   `json:"type" bson:"type"`
		Tags            []string        `json:"tags" bson:"tags"`
		Sample          string          `json:"sample" bson:"sample"`
		Test            string          `json:"test,omitempty" bson:"test"`
		Level           ChallengeLevel  `json:"level" bson:"level"`
		IsNew           bool            `json:"is_new" bson:"is_new"`
		CreatedAt       *time.Time      `json:"created_at" bson:"created_at"`
		UpdatedAt       *time.Time      `json:"updated_at" bson:"updated_at"`
		CreatedByID     string          `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedByName   string          `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByAvatar string          `json:"created_by_avatar,omitempty" bson:"created_by_avatar"`
	}
)
