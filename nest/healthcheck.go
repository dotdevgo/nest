package nest

import (
	"net/http"

	"github.com/defval/di"
)

// HealthCheck godoc
func HealthCheck() di.Option {
	return di.Options(
		di.Invoke(func(w *Kernel) {
			w.GET("/health-check", func(ctx Context) error {
				return ctx.JSON(http.StatusOK, Map{
					"Status": "OK",
				})
			})
		}),
	)
}
