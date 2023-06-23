package kernel

import (
	"dotdev/nest/pkg/kernel/extension"
	"dotdev/nest/pkg/nest"

	"github.com/defval/di"
)

// New godoc
func New() di.Option {
	nest.LoadEnv()

	return di.Options(
		extension.Bus(),
		extension.EventBus(),
		extension.HealthCheck(),
		extension.Validator(),
	)
}
