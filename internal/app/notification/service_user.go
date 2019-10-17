package notification

import (
	"fmt"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/event"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

func (s *Service) handleUserPasswordResetTokenCreated(ev event.Event) {
	var n types.ResetPasswordTokenCreated
	if err := ev.Data.Unmarshal(&n); err != nil {
		log.Errorf("failed to unmarshal notification, err: %v", err)
		return
	}
	subject := fmt.Sprintf("Password reset is requested")
	s.sendEmailNotification(subject, "user_password_reset_token_created.html", n, n.User.UserID)
}
