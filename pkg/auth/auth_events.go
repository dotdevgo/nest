package auth

import "dotdev/nest/pkg/user"

const (
	EventAuthSignUp     = "auth.sign_up"
	EventAuthConfirm    = "auth.confirm"
	EventAuthRestore    = "auth.restore"
	EventAuthResetToken = "auth.reset_token"
	EventAuthResetEmail = "auth.reset_email"
)

type (
	EventAuthGeneric struct {
		User *user.User
	}

	EventResetToken struct {
		User     *user.User
		Password string
	}
)
