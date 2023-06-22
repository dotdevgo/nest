package authcmd

import (
	"dotdev/nest/cmd/auth/controller"
	"dotdev/nest/pkg/nest"

	"github.com/defval/di"
)

// New godoc
func New() di.Option {
	return di.Options(
		nest.NewController(func() *controller.AuthController {
			return &controller.AuthController{}
		}),
	)
}
