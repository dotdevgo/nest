package crud

import (
	"database/sql"

	"dotdev/nest/pkg/nest"

	"gorm.io/gorm"
)

type Service[T IModel] struct {
	db *gorm.DB
	tx *gorm.DB
}

// NewService godoc
func NewService[T IModel](db *gorm.DB) *Service[T] {
	s := &Service[T]{db: db}
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
		return s.Tx().Save(data).Error
	}

	return s.Tx().Create(data).Error
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
	var stmt = s.Tx().Session(&gorm.Session{}) //NewDB: true

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}

// DB Godoc
func (s *Service[T]) DB() *gorm.DB {
	return s.db
}

// Tx Get current transaction
func (s *Service[T]) Tx() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

// Begin godoc
func (s *Service[T]) Begin(opts ...*sql.TxOptions) (*gorm.DB, error) {
	s.tx = s.db.Begin(opts...)

	defer func() {
		if r := recover(); r != nil && s.tx != nil {
			s.tx.Rollback()
		}
	}()

	if err := s.tx.Error; err != nil {
		return nil, err
	}

	return s.tx, nil
}

// Commit godoc
func (s *Service[T]) Commit() *gorm.DB {
	if s.tx == nil {
		return nil
	}

	tx := s.tx.Commit()
	s.tx = nil
	return tx
}

// Rollback godoc
func (s *Service[T]) Rollback() {
	if s.tx == nil {
		return
	}

	s.tx.Rollback()
	s.tx = nil
}

// Transaction godoc
func (s *Service[T]) Transaction(fc func(tx *gorm.DB) (err error), opts ...*sql.TxOptions) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		s.tx = tx
		err := fc(tx)
		s.tx = nil
		return err
	}, opts...)
}
