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

// // GetTableName godoc
// func GetTableName(db *gorm.DB, value interface{}) string {
// 	stmt := &gorm.Statement{DB: db}
// 	if err := stmt.Parse(value); err != nil {
// 		return err.Error()
// 	}
// 	return stmt.Schema.Table
// }
