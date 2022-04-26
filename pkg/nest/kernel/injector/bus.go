package injector

import (
	"dotdev/nest/pkg/nest/kernel/injector/factory"

	"github.com/goava/di"
	"github.com/mustafaturan/bus/v3"
)

// NewBus godoc
func NewBus() di.Option {
	return di.Options(
		di.Provide(func() *bus.Bus {
			return factory.NewBus()
		}),
	)
}
