package injector

import nest "github.com/dotdevgo/nest/pkg/nest"

// Router godoc
func NewApiGroup(e *nest.EchoWrapper) nest.ApiGroup {
	api := e.Group("/api")
	return api
}
