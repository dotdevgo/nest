package dataform

import (
	"dotdev.io/examples/app/dataform/handler/controller"
	"dotdev.io/examples/app/dataform/handler/graphql"
	"dotdev.io/examples/app/dataform/orm/repository"
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
