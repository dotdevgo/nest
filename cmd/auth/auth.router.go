package authcmd

import (
	"dotdev/nest/pkg/nest"

	"github.com/goava/di"
)

// NewRouter godoc
func NewRouter() di.Option {
	return di.Options(
		di.Provide(func() *AuthController {
			return &AuthController{}
		}, di.As(new(nest.Controller))),
		di.Provide(func() *UserController {
			return &UserController{}
		}, di.As(new(nest.Controller))),
	)
}
