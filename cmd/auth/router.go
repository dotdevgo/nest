package authcmd

import (
	"dotdev/nest/cmd/auth/controller"
	"dotdev/nest/pkg/nest"

	"github.com/goava/di"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Provide(func() *controller.AuthController {
			return &controller.AuthController{}
		}, di.As(new(nest.Controller))),
		// di.Provide(func() *controller.UserController {
		// 	return &controller.UserController{}
		// }, di.As(new(nest.Controller))),
	)
}
