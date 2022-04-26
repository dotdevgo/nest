package auth

import "dotdev/nest/pkg/user"

type IdentityDto struct {
	Identity string `json:"identity" form:"identity" validate:"required,min=3"`
}

type SignInDto struct {
	IdentityDto

	Password string `json:"password" form:"password" validate:"required"`
}

type SignUpDto struct {
	user.UserDto

	Password string `json:"password" form:"password" validate:"required,min=5"`
}
