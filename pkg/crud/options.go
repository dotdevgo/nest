package crud

import (
	"gorm.io/gorm"
)

type Option func(db *gorm.DB) *gorm.DB

func WithScope(funcs ...func(*gorm.DB) *gorm.DB) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(funcs...)
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
