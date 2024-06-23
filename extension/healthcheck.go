package extension

import (
	"dotdev/nest"
	"net/http"

	"github.com/defval/di"
)

// HealthCheck godoc
func HealthCheck() di.Option {
	return di.Options(
		di.Invoke(func(w *nest.Kernel) {
			w.GET("/health-check", func(ctx nest.Context) error {
				return ctx.String(http.StatusOK, "OK")
			})
		}),
	)
}
