package swagger

import (
	"dotdev/nest"

	"github.com/defval/di"
	swag "github.com/swaggo/echo-swagger"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Invoke(func(e *nest.Kernel) error {
			e.Echo.GET("/docs/*", swag.WrapHandler)

			return nil
		}),
	)
}
