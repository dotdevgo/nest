package handler

import (
	"dotdev.io/internal/app/dataform/dto"
	"dotdev.io/internal/app/dataform/entity"
	"dotdev.io/internal/app/dataform/repository"
	"dotdev.io/pkg/crud"
	paginator "dotdev.io/pkg/gorm-paginator"
	"dotdev.io/pkg/nest"
	"net/http"
)

// SaveFormTemplate godoc
func SaveFormTemplate(c nest.Context) interface{} {
	return func(s crud.Service) error {
		var input = new(dto.FormTemplate)
		if err := s.IsValid(c, input); err != nil {
			return nest.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var data = repository.NewFormTemplateEntity(input)
		if err := s.Save(&data); err != nil {
			return nest.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, data)
	}
}

// ListFormTemplate godoc
func ListFormTemplate(c nest.Context) interface{} {
	return func(s crud.Service) error {
		var result []entity.FormTemplate

		data, err := s.Paginate(
			&result,
			[]paginator.Option{
				paginator.WithRequest(c.Request()),
			},
			crud.WithRequest(c.Request()),
			//crud.WithScope(repository.FormTemplateScope(map[string]interface{}{
			//	"name": "test text",
			//})),
			//crud.WithCriteria(crud.Criteria{
			//	crud.CriteriaOption{Field: "name", Value: c.QueryParam("criteria[name]")},
			//}),
		)

		if err != nil {
			return nest.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, data)
	}
}
