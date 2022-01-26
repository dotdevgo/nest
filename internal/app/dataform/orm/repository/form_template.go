package repository

import (
	"dotdev.io/internal/app/dataform/orm/entity"
	"dotdev.io/pkg/goutils"
	"gorm.io/gorm"
)

// ScopeFormTemplate godoc
func ScopeFormTemplate(criteria map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if criteria["name"] != nil {
			db = db.Where("name = ?", criteria["name"])
		}
		return db
	}
}

// NewFormTemplate godoc
func NewFormTemplate(input interface{}) *entity.FormTemplate {
	data := new(entity.FormTemplate)
	goutils.Copy(data, input)
	return data
}
