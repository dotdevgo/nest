package handler

import (
	"dotdev.io/internal/app/dataform/dto"
	"dotdev.io/internal/app/dataform/entity"
	"dotdev.io/internal/app/dataform/factory"
	"dotdev.io/pkg/crud"
	"dotdev.io/pkg/nest"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SaveFormTemplate(c nest.Context) interface{} {
	return func(s *crud.Service) error {
		var input = new(dto.FormTemplate)
		if err := s.IsValid(c, input); err != nil {
			return err
		}

		var data = factory.NewFormTemplateEntity(input)
		if err := s.Create(&data); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, data)
	}
}

func ListFormTemplate(c nest.Context) interface{} {
	return func(s *crud.Service) error {
		var result []entity.FormTemplate

		paginator, err := s.Paginate(c, result)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, paginator)
	}
}

//func ListFormTemplate(c nest.Context) interface{} {
//	return func(s *crud.Service) error {
//		var result []entity.FormTemplate
//		s.Paginate(c, &result)
//		//if err := s.Paginate(c, &result); err != nil {
//		//	return err
//		//}
//		return c.JSON(http.StatusOK, result)
//	}
//}

//if err := c.Bind(input); err != nil {
//	return c.String(http.StatusBadRequest, err.Error())
//}
//
//if err := validate.Struct(input); err != nil {
//	return utils.ValidationError(c, err)
//}

//data := new(entity.FormTemplate)
//
//copiers := copy.New()
//copiers.Copy(data, input)
//func ListFormTemplate(c nest.Context) error {
//	return c.JSON(http.StatusOK, map[string]interface{}{"status": "form-template"})
//}
