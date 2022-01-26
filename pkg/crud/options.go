package crud

import (
	"gorm.io/gorm"
)

type Option func(db *gorm.DB) *gorm.DB

// WithScope godoc
func WithScope(funcs ...func(*gorm.DB) *gorm.DB) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(funcs...)
	}
}

// WithPreload godoc
func WithPreload(model string, args ...interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(model, args...)
	}
}

//var sqlOperator = "="
//var sqlValue = val
//var parts = strings.Split(val.(string), "||")
//
//if len(parts) == 2 {
//	sqlValue = parts[1]
//
//	switch parts[0] {
//	case "$eq":
//		sqlOperator = "="
//		break
//	case "$cont":
//		sqlOperator = "LIKE"
//		sqlValue = "%" + parts[1] + "%"
//		break
//	}
//}
//db = db.Where(name+" "+sqlOperator+" ?", sqlValue)
