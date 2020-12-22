package dataform

import (
	"dotdev.io/internal/app/dataform/entity"
	"dotdev.io/internal/app/dataform/handler"
	"dotdev.io/pkg/nest"
	"dotdev.io/pkg/nest/provider"
)

var OrmConfig = &provider.OrmConfig{
	Entities: []interface{}{
		&entity.FormTemplate{},
	},
}

func Router(e *nest.EchoWrapper) {
	e.GET("/form-template", e.HandlerFn(handler.ListFormTemplate))
	e.POST("/form-template", e.HandlerFn(handler.SaveFormTemplate))
	e.PUT("/form-template/:id", e.HandlerFn(handler.SaveFormTemplate))
}
