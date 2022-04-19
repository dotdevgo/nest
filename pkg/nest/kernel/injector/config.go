package injector

import (
	"github.com/dotdevgo/nest/pkg/goutils"
	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/goava/di"
)

// Config godoc
func Config() di.Option {
	return di.Options(
		di.Provide(func(w *nest.Kernel) nest.Config {
			config, err := nest.GetConfig()
			goutils.NoErrorOrFatal(err)

			w.Config = config

			return config
		}),
	)
}
