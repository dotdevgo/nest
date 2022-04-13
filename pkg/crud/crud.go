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
func (s *Service[T]) IsValid(c nest.Context, input interface{}) error {
	if err := c.Bind(input); err != nil {
		return err
	}

	if err := c.Validate(input); err != nil {
		return err
	}

	return nil
}

// Save godoc
func (s *Service[T]) Save(data T) error {
	if data.GetID() > 0 || data.GetUUID() != "" {
		return s.DB.Save(data).Error
	}

	return s.DB.Create(data).Error
}

// Find godoc
func (s *Service[T]) Find(result T, id interface{}, options ...Option) error {
	stmt := s.NewStmt(options...)

	return stmt.Scopes(ScopeById(result, id)).First(result).Error
}

// GetMany godoc
func (s *Service[T]) GetMany(result interface{}, options ...Option) error {
	var stmt = s.NewStmt(options...)

	return stmt.Find(result).Error
}

// NewStmt godoc
func (s *Service[T]) NewStmt(options ...Option) *gorm.DB {
	var stmt = s.DB.Session(&gorm.Session{}) //NewDB: true

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}
