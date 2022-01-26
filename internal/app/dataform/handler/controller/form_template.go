package controller

import (
	"dotdev.io/internal/app/dataform/dto"
	"dotdev.io/internal/app/dataform/orm/entity"
	"dotdev.io/internal/app/dataform/orm/repository"
	"dotdev.io/pkg/crud"
	paginator "dotdev.io/pkg/gorm-paginator"
	"dotdev.io/pkg/nest"
	"dotdev.io/pkg/nest/httpkernel"
	"net/http"
)

// FormTemplateController godoc
type FormTemplateController struct {
	httpkernel.Controller

	Crud *crud.Service
}

// Save godoc
func (c *FormTemplateController) Save(ctx nest.Context) error {
	var input = new(dto.FormTemplateDto)
	if err := c.Crud.IsValid(ctx, input); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var data = repository.NewFormTemplate(input)
	if err := c.Crud.Save(data); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, data)
}

// List godoc
func (c *FormTemplateController) List(ctx nest.Context) error {
	var result []entity.FormTemplate

	data, err := c.Crud.Paginate(
		&result,
		[]paginator.Option{
			paginator.WithRequest(ctx.Request()),
		},
		crud.WithRequest(ctx.Request()),
	)

	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, data)
}

// ListFormTemplate godoc
//func ListFormTemplate(c nest.Context) interface{} {
//	return func(s *crud.Service) error {
//		var result []entity.FormTemplateDto
//
//		data, err := s.Paginate(
//			&result,
//			[]paginator.Option{
//				paginator.WithRequest(c.Request()),
//			},
//			crud.WithRequest(c.Request()),
//			//crud.WithScope(repository.FormTemplateScope(map[string]interface{}{
//			//	"name": "test text",
//			//})),
//			//crud.WithCriteria(crud.Criteria{
//			//	crud.CriteriaOption{Field: "name", Value: c.QueryParam("criteria[name]")},
//			//}),
//		)
//
//		if err != nil {
//			return nest.NewHTTPError(http.StatusBadRequest, err.Error())
//		}
//
//		return c.JSON(http.StatusOK, data)
//	}
//}
