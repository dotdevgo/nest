package httpkernel

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

//// CrudController godoc
//type CrudController struct {
//	di.Inject
//
//	Crud *crud.Service
//}
//
//// NewController creates controller.
//func NewCrudController(crud *crud.Service) *CrudController {
//	return &CrudController{Crud: crud}
//}
