package dataform

import (
	"dotdev.io/internal/app/dataform/handler/controller"
	"dotdev.io/internal/app/dataform/handler/graphql"
	"dotdev.io/internal/app/dataform/orm/repository"
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
