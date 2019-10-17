package user

type (
	GenerateResetPasswordTokenRequest struct {
		Email string `json:"email"`
	}

	ResetPasswordRequest struct {
		Token       string `json:"token" validate:"required"`
		NewPassword string `json:"new_password" validate:"required"`
	}
)
