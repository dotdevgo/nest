package extension

import (
	"dotdev/nest/pkg/logger"
	"dotdev/nest/pkg/nest"

	"github.com/defval/di"
)

// Config godoc
func Config() di.Option {
	return di.Options(
		di.Provide(func(w *nest.Kernel) nest.Config {
			config, err := nest.GetConfig()
			logger.FatalOnError(err)

			w.Config = config

			return config
		}),
	)
}
