package controller

import (
	"dotdev.io/examples/app/dataform/dto"
	"dotdev.io/examples/app/dataform/orm/entity"
	"dotdev.io/examples/app/dataform/orm/repository"
	"dotdev.io/pkg/crud"
	"dotdev.io/pkg/nest"
	"dotdev.io/pkg/nest/kernel"
	"net/http"
)

// FormTemplateController godoc
type FormTemplateController struct {
	kernel.Controller
	Crud *crud.Service
	Repo *repository.FormTemplateRepo
}

// Save godoc
func (c *FormTemplateController) Save(ctx nest.Context) error {
	var input = new(dto.FormTemplateDto)

	if err := c.Crud.IsValid(ctx, input); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	data, err := c.Repo.FindOrNew(input)
	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Crud.Save(data); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, data)
}

// List godoc
func (c *FormTemplateController) List(ctx nest.Context) error {
	var result []entity.FormTemplate

	data, err := c.Crud.Paginate(
		&result,
		crud.WithRequestCursor(ctx.Request()),
		crud.WithRequest(ctx.Request()),
	)

	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, data)
}
