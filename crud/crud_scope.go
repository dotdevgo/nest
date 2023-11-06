package crud

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ScopeOrderBy godoc
func ScopeOrderBy(column string, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", column, strings.ToUpper(order)))
	}
}

// ScopeById godoc
//func ScopeById[T any](result T, id interface{}) func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//		table := GetTableName(db, result)
//
//		uid, ok := id.(string)
//		if ok {
//			_, err := uuid.Parse(uid)
//			if err == nil {
//				return db.Where(table+".id = ?", id)
//			}
//		}
//
//		return db.Where(table+".pk = ?", id)
//	}
//}
