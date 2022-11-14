package kernel

import (
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/nest/kernel/injector"

	"github.com/goava/di"
)

// New godoc
func New() di.Option {
	nest.LoadEnv()

	return di.Options(
		injector.Config(),
		injector.EventBus(),
		injector.NewBus(),
		injector.Validator(),
		injector.I18n(),
		di.Provide(injector.NewSecureGroup),
	)
}
