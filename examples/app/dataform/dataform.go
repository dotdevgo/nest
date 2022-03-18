package dataform

import (
	"github.com/dotdevgo/nest/examples/app/dataform/handler/controller"
	"github.com/dotdevgo/nest/examples/app/dataform/orm/entity"
	nest "github.com/dotdevgo/nest/pkg/core"
	"github.com/goava/di"
	"gorm.io/gorm"
)

// Router godoc
func Router(api nest.ApiGroup, formTemplate *controller.FormTemplateController) {
	e := api.(*nest.Group)
	e.GET("/form-template", formTemplate.List)
	e.POST("/form-template", formTemplate.Save)
	e.PUT("/form-template/:id", formTemplate.Save)
}

// Provider godoc
func Provider() di.Option {
	return di.Options(
		di.Provide(NewFormTemplateCtrl),
		di.Provide(NewFormTemplateResolver),
		di.Provide(NewFormTemplateRepo),
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&entity.FormTemplate{})
		}),
	)
}
