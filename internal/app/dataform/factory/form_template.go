package factory

import (
	"dotdev.io/internal/app/dataform/entity"
	"github.com/gotidy/copy"
)

func NewFormTemplateEntity(input interface{}) *entity.FormTemplate {
	data := new(entity.FormTemplate)

	copiers := copy.New()
	copiers.Copy(data, input)

	return data
}
