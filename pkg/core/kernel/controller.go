package kernel

import (
	"github.com/goava/di"
)

type Controller struct {
	di.Inject
}

// NewController creates controller.
func NewController() *Controller {
	return &Controller{}
}
