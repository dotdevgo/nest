package kernel

import (
	"github.com/dotdevgo/nest/pkg/core/kernel/injector"
	"github.com/goava/di"
)

// Provider godoc
func Provider() di.Option {
	return di.Options(
		// injector.OrmDefault(),
		injector.Validator(),
		injector.Crud(),
		injector.EventBus(),
		di.Provide(injector.Router),
	)
}
