package extension

import (
	"github.com/mustafaturan/monoton/v2"
	"github.com/mustafaturan/monoton/v2/sequencer"

	"github.com/defval/di"
	"github.com/mustafaturan/bus/v3"
)

// Bus godoc
func Bus() di.Option {
	return di.Options(
		di.Provide(func() *bus.Bus {
			// configure id generator (it doesn't have to be monoton)
			node := uint64(1)
			initialTime := uint64(1577865600000) // set 2020-01-01 PST as initial time
			m, err := monoton.New(sequencer.NewMillisecond(), node, initialTime)
			if err != nil {
				panic(err)
			}

			// init an id generator
			var idGenerator bus.Next = m.Next

			// create a new bus instance
			b, err := bus.NewBus(idGenerator)
			if err != nil {
				panic(err)
			}

			return b
		}),
	)
}
