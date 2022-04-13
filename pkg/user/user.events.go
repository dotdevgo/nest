package user

const (
	EventUserSignUp     = "user.sign_up"
	EventUserConfirm    = "user.confirm"
	EventUserRestore    = "user.restore"
	EventUserResetToken = "user.reset_token"
)

type (
	EventResetToken struct {
		User     *User
		Password string
	}
)
