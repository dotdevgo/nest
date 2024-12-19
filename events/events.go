package events

import (
	eventBus "github.com/asaskevich/EventBus"
	"github.com/defval/di"
)

func New() di.Option {
	return di.Options(
		di.Provide(func() eventBus.Bus {
			return eventBus.New()
		}),
	)
}
