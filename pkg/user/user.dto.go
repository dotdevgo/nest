package user

import "github.com/dotdevgo/nest/pkg/crud"

// UserDto godoc
type UserDto struct {
	crud.Model

	Email       string  `json:"email" form:"email" validate:"required,email,uniqueEmail"`
	Username    string  `json:"username" form:"username" validate:"required,uniqueUsername,min=5"`
	DisplayName *string `json:"displayName" form:"displayName" validate:"omitempty,min=5"`

	Bio string `json:"bio" form:"bio"`
}
