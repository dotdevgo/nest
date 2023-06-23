package orm

import (
	"github.com/defval/di"
	"gorm.io/gorm"
)

// Mirgrate godoc
func Mirgrate(entity any) di.Option {
	return di.Options(
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(entity)
		}),
	)
}
