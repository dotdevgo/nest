package kernel

import (
	"github.com/dotdevgo/gosymfony/pkg/nest"
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
