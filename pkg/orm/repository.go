package orm

import (
	"dotdev/nest/pkg/crud"

	"github.com/defval/di"
	"gorm.io/gorm"
)

// Option godoc
type Option func(db *gorm.DB) *gorm.DB

// Repository godoc
type Repository[T crud.Model] struct {
	di.Inject
	*gorm.DB
}

// CreateQueryBuilder godoc
func (r Repository[T]) CreateQueryBuilder(options ...Option) *gorm.DB {
	var stmt = r.Session(&gorm.Session{})

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}

// Add godoc
func (r Repository[T]) Add(data T) error {
	if data.GetID() != "" {
		return r.Save(data).Error
	}

	return r.Create(data).Error
}
