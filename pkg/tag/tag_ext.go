package tag

import (
	"dotdev/nest/pkg/crud"

	"github.com/defval/di"
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
func NewTagCrud(c *crud.Crud[*Tag]) *TagCrud {
	return &TagCrud{
		Crud: c,
	}
}
