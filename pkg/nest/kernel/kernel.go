package kernel

import (
	"dotdev/nest/pkg/nest/kernel/injector"

	"github.com/goava/di"
)

// Provider godoc
func New() di.Option {
	return di.Options(
		injector.Config(),
		injector.EventBus(),
		injector.NewBus(),
		injector.Validator(),
		injector.I18n(),
		di.Provide(injector.NewApiGroup),
	)
}
