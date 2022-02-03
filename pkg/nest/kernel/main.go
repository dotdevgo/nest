package kernel

import (
	"dotdev.io/pkg/nest/kernel/injector"
	"github.com/goava/di"
)

func Provider() di.Option {
	return di.Options(
		injector.Validator(),
		injector.Crud(),
		injector.EventBus(),
		di.Provide(NewRouterGroupApi),
	)
}

