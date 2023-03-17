package injector

import (
	"dotdev/nest/pkg/nest"

	"github.com/goava/di"
)

// Router godoc
func Router() di.Option {
	return di.Options(
		di.Provide(secureGroup),
	)
}

func secureGroup(e *nest.Kernel) nest.SecureGroup {
	g := e.Group("")
	return g
}
