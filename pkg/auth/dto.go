package auth

type IdentityDto struct {
	Identity string `json:"identity" form:"identity" validate:"required,min=3"`
}

type SignInDto struct {
	IdentityDto

	Password string `json:"password" form:"password" validate:"required"`
}

type SignUpDto struct {
	Email    string `json:"email" form:"email" validate:"required,email,uniqueEmail"`
	Username string `json:"username" form:"username" validate:"required,uniqueUsername,min=5"`
	Password string `json:"password" form:"password" validate:"required,min=5"`
}

type ChangePasswordDto struct {
	Password        string `json:"password" form:"password" validate:"required,min=5"`
	NewPassword     string `json:"newPassword" form:"newPassword" validate:"required,min=5"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required,min=5,eqfield=NewPassword"`
}
