package crud

import (
	"github.com/dotdevgo/nest/pkg/nest"
	"gorm.io/gorm"
)

type Service[T IModel] struct {
	DB *gorm.DB
}

// NewService godoc
func NewService[T IModel](db *gorm.DB) *Service[T] {
	s := &Service[T]{DB: db}
	return s
	// .Session(&gorm.Session{NewDB: true})
}

// IsValid godoc
func (s *Service[T]) IsValid(ctx nest.Context, input interface{}) error {
	if err := ctx.Bind(input); err != nil {
		return err
	}

	if err := ctx.Validate(input); err != nil {
		return err
	}

	return nil
}

// Save godoc
func (s *Service[T]) Flush(data T) error {
	if data.GetPk() > 0 || data.GetID() != "" {
		return s.DB.Save(data).Error
	}

	return s.DB.Create(data).Error
}

// Find godoc
func (s *Service[T]) Find(result T, id interface{}, options ...Option) error {
	stmt := s.Stmt(options...)

	return stmt.Scopes(ScopeById(result, id)).First(result).Error
}

// GetMany godoc
func (s *Service[T]) GetMany(result interface{}, options ...Option) error {
	var stmt = s.Stmt(options...)

	return stmt.Find(result).Error
}

// Stmt godoc
func (s *Service[T]) Stmt(options ...Option) *gorm.DB {
	var stmt = s.DB.Session(&gorm.Session{}) //NewDB: true

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}
