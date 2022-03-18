package controller

import (
	"net/http"

	"github.com/dotdevgo/nest/examples/app/hltv/dto"
	"github.com/dotdevgo/nest/examples/app/hltv/orm/entity"
	"github.com/dotdevgo/nest/examples/app/hltv/orm/repository"
	nest "github.com/dotdevgo/nest/pkg/core"
	"github.com/dotdevgo/nest/pkg/core/kernel"
	"github.com/dotdevgo/nest/pkg/crud"
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
