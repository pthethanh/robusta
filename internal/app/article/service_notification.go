package article

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/email"
	"github.com/pthethanh/robusta/internal/pkg/log"
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
	err = s.mailer.Send(context.Background(), email.Email{
		To:      []string{user.Email},
		Subject: fmt.Sprintf("%s commented on your post - %s", c.CreatedByName, a.Title),
		Body:    fmt.Sprintf(`Article: <a href="%s">%s</a><br>Comment: %s<br>Goway`, a.GetLink(), a.Title, c.Content),
	})
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}
	log.Debugf("notification: comment_added is already sent")
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
	err = s.mailer.Send(context.Background(), email.Email{
		To:      []string{user.Email},
		Subject: fmt.Sprintf("%s %s your post - %s", r.CreatedByName, r.Type, a.Title),
		Body:    fmt.Sprintf(`Article: <a href="%s">%s</a><br>Goway`, a.GetLink(), a.Title),
	})
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}
	log.Debugf("notification: reaction  %s is already sent", r.Type)
	return nil
}
