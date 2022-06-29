package user

import "dotdev/nest/pkg/crud"

// UserDto godoc
type UserDto struct {
	crud.Model
	crud.Attributes

	Email       string `json:"email" form:"email" validate:"omitempty,email,uniqueEmail"`
	Username    string `json:"username" form:"username" validate:"omitempty,uniqueUsername,min=5"`
	DisplayName string `json:"displayName" form:"displayName" validate:"omitempty,min=5"`

	Bio string `json:"bio" form:"bio"`
}
