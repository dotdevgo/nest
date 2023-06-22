package user

import (
	"fmt"

	"gorm.io/gorm"
)

// ScopeByIdentity godoc
func ScopeByIdentity(identity string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		sql := fmt.Sprintf("(%s.email = ? OR %s.username = ?)", DBTableUsers, DBTableUsers)

		return db.Where(sql, identity, identity)
	}
}
