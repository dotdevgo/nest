package dataform

import (
	"github.com/dotdevgo/nest/examples/app/dataform/handler/controller"
	"github.com/dotdevgo/nest/examples/app/dataform/handler/graphql"
	"github.com/dotdevgo/nest/examples/app/dataform/orm/repository"
)

// NewController creates controller.
func NewFormTemplateCtrl() *controller.FormTemplateController {
	return &controller.FormTemplateController{}
}

// NewFormTemplateResolver godoc
func NewFormTemplateResolver() *graphql.FormTemplateResolver {
	return &graphql.FormTemplateResolver{}
}

// NewFormTemplateRepo godoc
func NewFormTemplateRepo() *repository.FormTemplateRepo {
	return &repository.FormTemplateRepo{}
}
