package repository

import (
	"dotdev.io/internal/app/dataform/entity"
	"dotdev.io/pkg/goutils"
	"gorm.io/gorm"
)

// FormTemplateScope godoc
func FormTemplateScope(criteria map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if criteria["name"] != nil {
			db = db.Where("name = ?", criteria["name"])
		}
		return db
	}
}

// NewFormTemplateEntity godoc
func NewFormTemplateEntity(input interface{}) entity.FormTemplate {
	data := new(entity.FormTemplate)
	goutils.Copy(data, input)
	return *data
}
