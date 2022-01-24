package entity

import (
	"dotdev.io/pkg/crud"
)

type FormTemplate struct {
	crud.Model
	Name string `gorm:"type:string" json:"name"`
}
