package auth

import "github.com/dotdevgo/nest/pkg/user"

type RestoreDto struct {
	Identity string `json:"identity" form:"identity" validate:"required,min=3"`
}

type SignInDto struct {
	Identity string `json:"identity" form:"identity" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type SignUpDto struct {
	user.UserDto

	Password string `json:"password" form:"password" validate:"required,min=5"`
}