package events

import (
	events "github.com/asaskevich/EventBus"
	"github.com/defval/di"
)

func New() di.Option {
	return di.Options(
		di.Provide(func() events.Bus {
			return events.New()
		}),
	)
}
