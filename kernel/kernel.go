package kernel

import (
	"dotdev/nest/kernel/extension"
	"dotdev/nest/nest"

	"github.com/defval/di"
)

// New godoc
func New() di.Option {
	nest.LoadEnv()

	return di.Options(
		extension.HealthCheck(),
		extension.Validator(),
	)
}
