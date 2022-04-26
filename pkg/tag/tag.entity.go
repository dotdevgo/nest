package tag

import "dotdev/nest/pkg/crud"

const (
	DBTableTags = "tags"
)

type Tag struct {
	crud.Model

	Name string `json:"name" form:"name" gorm:"not null"`
}
