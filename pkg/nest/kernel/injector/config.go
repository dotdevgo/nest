package injector

import (
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/utils"

	"github.com/goava/di"
)

// Config godoc
func Config() di.Option {
	return di.Options(
		di.Provide(func(w *nest.Kernel) nest.Config {
			config, err := nest.GetConfig()
			utils.NoErrorOrFatal(err)

			w.Config = config

			return config
		}),
	)
}
