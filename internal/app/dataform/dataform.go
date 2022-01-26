package dataform

import (
	"dotdev.io/internal/app/dataform/handler/controller"
	"dotdev.io/internal/app/dataform/orm/entity"
	"dotdev.io/pkg/nest"
	"github.com/goava/di"
	"gorm.io/gorm"
)

// Router godoc
func Router(e *nest.EchoWrapper, formTemplate *controller.FormTemplateController) {
	e.GET("/form-template", formTemplate.List)
	e.POST("/form-template", formTemplate.Save)
	e.PUT("/form-template/:id", formTemplate.Save)
}

// Provider godoc
func Provider() di.Option {
	return di.Options(
		di.Provide(newFormTemplateCtrl),
		di.Provide(newFormTemplateResolver),
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&entity.FormTemplate{})
		}),
	)
}
