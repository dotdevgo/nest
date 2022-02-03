//go:generate metatag

package entity

import (
	"github.com/dotdevgo/gosymfony/pkg/crud"
)

const TableFormTemplate = "form_templates"

// FormTemplate godoc
type FormTemplate struct {
	crud.Model
	crud.Timestampable
	crud.SoftDeleteable

	Name string `gorm:"type:string" json:"name" meta:"getter;setter;"`
}

func (FormTemplate) IsRecord() {}
