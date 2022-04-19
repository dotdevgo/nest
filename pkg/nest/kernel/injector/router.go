package injector

import nest "github.com/dotdevgo/nest/pkg/nest"

// Router godoc
func NewApiGroup(e *nest.Kernel) nest.ApiGroup {
	api := e.Group("/api")
	return api
}
