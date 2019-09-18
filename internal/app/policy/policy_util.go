package policy

import (
	"context"
	"fmt"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

// GroupSubject return policy name of group
func GroupSubject(name string) Subject {
	return Subject(fmt.Sprintf("group_%s", name))
}

// UserSubject return policy name of the user
func UserSubject(name string) Subject {
	return Subject(fmt.Sprintf("user_%s", name))
}

// ArticleObject return name of article object
func ArticleObject(articleID string) Object {
	return Object(fmt.Sprintf("article_%s", articleID))
}

// TutorialObject return name of tutorial object
func TutorialObject(tutorialID string) Object {
	return Object(fmt.Sprintf("tutorial_%s", tutorialID))
}

// CommentObject return name of comment object
func CommentObject(id string) Object {
	return Object(fmt.Sprintf("comment_%s", id))
}

// SolutionObject return name of comment object
func SolutionObject(id string) Object {
	return Object(fmt.Sprintf("solution_%s", id))
}

func ChallengeObject(id string) Object {
	return Object(fmt.Sprintf("challenge_%s", id))
}

func FolderObject(id string) Object {
	return Object(fmt.Sprintf("folder_%s", id))
}

// IsCurrentUserAllowed is a util to check if the current user is allowed to do something
// the user context is expected to be existed in the given context
func IsCurrentUserAllowed(ctx context.Context, srv interface {
	IsAllowed(ctx context.Context, sub Subject, obj Object, act Action) bool
}, obj Object, act Action) error {
	user := auth.FromContext(ctx)
	if user == nil {
		return ErrNotAllowed
	}
	isAllowed := srv.IsAllowed(ctx, UserSubject(user.UserID), obj, act)
	if !isAllowed {
		log.WithContext(ctx).WithFields(log.Fields{"user_id": user.UserID, "action": act, "obj": obj}).Errorf("the user is not authorized to do the action")
		return ErrNotAllowed
	}
	return nil
}
