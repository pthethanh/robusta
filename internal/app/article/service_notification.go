package article

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/email"
)

func (s *Service) sendCommentCreatedNotification(c types.Comment) error {
	a, err := s.FindByID(context.Background(), c.Target)
	if err != nil {
		return errors.Wrap(err, "failed to find article")
	}
	// if it's the owner's comment, don't  need to send notification
	if a.CreatedByID == c.CreatedByID {
		return nil
	}
	// send email to the owner
	user, err := s.user.FindByUserID(context.Background(), a.CreatedByID)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}
	// TODO improve email content by using beautiful HTML template
	m := email.Email{
		To:      []string{user.Email},
		Subject: fmt.Sprintf("%s commented on your post - %s", c.CreatedByName, a.Title),
		Body:    fmt.Sprintf(`Article: <a href="%s">%s</a><br>Comment: %s<br>Goway`, a.GetLink(), a.Title, c.Content),
	}
	noti, err := types.NewNotification(types.NotificationTypeEmail, m, time.Now())
	if err != nil {
		return errors.Wrap(err, "failed to create notification")
	}
	s.notifier.Notify(context.Background(), noti)
	return nil
}

func (s *Service) sendReactionNotification(r types.Reaction) error {
	a, err := s.FindByID(context.Background(), r.TargetID)
	if err != nil {
		return errors.Wrap(err, "failed to find article")
	}
	// if it's the owner's reaction, don't  need to send notification
	if a.CreatedByID == r.CreatedByID {
		return nil
	}
	// send email to the owner
	user, err := s.user.FindByUserID(context.Background(), a.CreatedByID)
	if err != nil {
		return errors.Wrap(err, "failed to find user")
	}
	// TODO improve email content by using beautiful HTML template
	m := email.Email{
		To:      []string{user.Email},
		Subject: fmt.Sprintf("%s %s your post - %s", r.CreatedByName, r.Type, a.Title),
		Body:    fmt.Sprintf(`Article: <a href="%s">%s</a><br>Goway`, a.GetLink(), a.Title),
	}
	noti, err := types.NewNotification(types.NotificationTypeEmail, m, time.Now())
	if err != nil {
		return errors.Wrap(err, "failed to create notification")
	}
	s.notifier.Notify(context.Background(), noti)
	return nil
}
