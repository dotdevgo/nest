package injector

import nest "github.com/dotdevgo/nest/pkg/core"

// Router godoc
func Router(e *nest.EchoWrapper) nest.ApiGroup {
	api := e.Group("/api")
	//api.Use(JwtMiddleware())
	return api
}
