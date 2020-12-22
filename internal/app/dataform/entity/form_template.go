package entity

import (
	"dotdev.io/pkg/crud"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FormTemplate struct {
	crud.Model
	Name string `gorm:"type:string" json:"name"`
}

func (u *FormTemplate) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New().String()

	return
}
