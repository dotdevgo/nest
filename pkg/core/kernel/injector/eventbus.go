package injector

import (
	evbus "github.com/asaskevich/EventBus"
	"github.com/goava/di"
)

func EventBus() di.Option {
	return di.Options(
		di.Provide(func() evbus.Bus {
			return evbus.New()
		}),
	)
}
