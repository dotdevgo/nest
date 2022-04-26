package tag

import (
	"dotdev/nest/pkg/crud"

	"github.com/goava/di"
	"gorm.io/gorm"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&Tag{})
		}),
		di.Provide(crud.NewService[*Tag]),
		di.Provide(NewTagCrud),
	)
}

// NewTagCrud godoc
func NewTagCrud(c *crud.Service[*Tag]) *TagCrud {
	return &TagCrud{
		Service: c,
	}
}
