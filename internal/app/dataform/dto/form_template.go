package dto

import (
	"dotdev.io/internal/app/dataform/orm/entity"
	paginator "dotdev.io/pkg/gorm-paginator"
)

// FormTemplateDto godoc
type FormTemplateDto struct {
	UUID string `form:"id" json:"id" param:"id" gqlgen:"id"`
	Name string `form:"name" json:"name" gqlgen:"name" validate:"required"`
}

// FormTemplatePaginator godoc
type FormTemplatePaginator struct {
	Cursor  *paginator.Result  `json:"cursor"`
	Records []*entity.FormTemplate `json:"records"`
}
