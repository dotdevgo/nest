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
