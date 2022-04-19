package auth

import "github.com/dotdevgo/nest/pkg/user"

const (
	EventUserSignUp     = "user.sign_up"
	EventUserConfirm    = "user.confirm"
	EventUserRestore    = "user.restore"
	EventUserResetToken = "user.reset_token"
)

type (
	EventResetToken struct {
		User     user.User
		Password string
	}
)
