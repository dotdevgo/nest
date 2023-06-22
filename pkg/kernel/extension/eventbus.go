package extension

import (
	evbus "github.com/asaskevich/EventBus"
	"github.com/defval/di"
)

func EventBus() di.Option {
	return di.Options(
		di.Provide(func() evbus.Bus {
			return evbus.New()
		}),
	)
}
