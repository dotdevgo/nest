package auth

import "dotdev/nest/pkg/user"

const (
	EventUserSignUp     = "user.sign_up"
	EventUserConfirm    = "user.confirm"
	EventUserRestore    = "user.restore"
	EventUserResetToken = "user.reset_token"
	EventUserResetEmail = "user.reset_email"
)

type (
	EventResetToken struct {
		User     user.User
		Password string
	}
)
