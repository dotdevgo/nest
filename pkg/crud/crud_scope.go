package crud

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

// ScopeOrderBy godoc
func ScopeOrderBy(column string, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", column, strings.ToUpper(order)))
	}
}

// ScopeById godoc
func ScopeById(result interface{}, id interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		table := GetTableName(db, result)
		str, ok := id.(string)
		if ok {
			_, err := uuid.Parse(str)
			if err == nil {
				return db.Where(table+".id = ?", id)
			}
		}
		return db.Where(table+".pk = ?", id)
	}
}
