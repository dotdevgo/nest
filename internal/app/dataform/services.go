package dataform

import (
	"dotdev.io/internal/app/dataform/handler/controller"
	"dotdev.io/internal/app/dataform/handler/graphql"
)

// NewController creates controller.
func newFormTemplateCtrl() *controller.FormTemplateController {
	return &controller.FormTemplateController{}
}

// newFormTemplateResolver godoc
func newFormTemplateResolver() *graphql.FormTemplateResolver {
	return &graphql.FormTemplateResolver{}
}
