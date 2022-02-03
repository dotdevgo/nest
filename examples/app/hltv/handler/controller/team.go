package controller

import (
	"github.com/dotdevgo/gosymfony/examples/app/hltv/dto"
	"github.com/dotdevgo/gosymfony/examples/app/hltv/orm/entity"
	"github.com/dotdevgo/gosymfony/examples/app/hltv/orm/repository"
	"github.com/dotdevgo/gosymfony/pkg/crud"
	"github.com/dotdevgo/gosymfony/pkg/nest"
	"github.com/dotdevgo/gosymfony/pkg/nest/kernel"
	"net/http"
)

// TeamController godoc
type TeamController struct {
	kernel.Controller
	Crud *crud.Service
	Repo *repository.TeamRepo
}

// Save godoc
func (c *TeamController) Save(ctx nest.Context) error {
	var input = new(dto.TeamDto)

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
func (c *TeamController) List(ctx nest.Context) error {
	var result []entity.Team

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
