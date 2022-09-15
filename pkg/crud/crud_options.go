package crud

import (
	"context"

	"gorm.io/gorm"
)

type Option func(db *gorm.DB) *gorm.DB

// WithScope godoc
func WithScope(funcs ...func(*gorm.DB) *gorm.DB) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(funcs...)
	}
}

// WithSelect godoc
func WithSelect(query interface{}, args ...interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select(query, args...)
	}
}

// WithPreload godoc
func WithPreload(model string, args ...interface{}) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(model, args...)
	}
}

// WithContext godoc
func WithContext(ctx context.Context) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.WithContext(ctx)
	}
}
