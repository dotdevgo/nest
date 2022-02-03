package kernel

import (
	"dotdev.io/pkg/nest"
	"github.com/goava/di"
)

type Controller struct {
	di.Inject
}

// NewController creates controller.
func NewController() *Controller {
	return &Controller{}
}

// NewRouterGroupApi godoc
func NewRouterGroupApi(e *nest.EchoWrapper) nest.ApiGroup {
	api := e.Group("/api")
	//api.Use(JwtMiddleware())
	return api
}
