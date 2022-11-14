package crud

import (
	"database/sql"

	"dotdev/nest/pkg/nest"

	"gorm.io/gorm"
)

type Crud[T IModel] struct {
	db *gorm.DB
	tx *gorm.DB
}

// NewService godoc
func NewService[T IModel](db *gorm.DB) *Crud[T] {
	s := &Crud[T]{db: db}
	return s
	// .Session(&gorm.Session{NewDB: true})
}

// IsValid godoc
func (s *Crud[T]) IsValid(ctx nest.Context, input interface{}) error {
	if err := ctx.Bind(input); err != nil {
		return err
	}

	if err := ctx.Validate(input); err != nil {
		return err
	}

	return nil
}

// Flush godoc
func (s *Crud[T]) Flush(data T) error {
	if data.GetPk() > 0 || data.GetID() != "" {
		return s.Tx().Save(data).Error
	}

	return s.Tx().Create(data).Error
}

// Find godoc
func (s *Crud[T]) Find(result T, id interface{}, options ...Option) error {
	stmt := s.Stmt(options...)

	return stmt.Scopes(ScopeById(result, id)).First(result).Error
}

// GetMany godoc
func (s *Crud[T]) GetMany(result interface{}, options ...Option) error {
	var stmt = s.Stmt(options...)

	return stmt.Find(result).Error
}

// Stmt godoc
func (s *Crud[T]) Stmt(options ...Option) *gorm.DB {
	var stmt = s.Tx().Session(&gorm.Session{}) //NewDB: true

	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}

// DB Godoc
func (s *Crud[T]) DB() *gorm.DB {
	return s.db
}

// Tx Get current transaction
func (s *Crud[T]) Tx() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

// Begin godoc
func (s *Crud[T]) Begin(opts ...*sql.TxOptions) (*gorm.DB, error) {
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
func (s *Crud[T]) Commit() *gorm.DB {
	if s.tx == nil {
		return nil
	}

	tx := s.tx.Commit()
	s.tx = nil
	return tx
}

// Rollback godoc
func (s *Crud[T]) Rollback() {
	if s.tx == nil {
		return
	}

	s.tx.Rollback()
	s.tx = nil
}

// Transaction godoc
func (s *Crud[T]) Transaction(fc func(tx *gorm.DB) (err error), opts ...*sql.TxOptions) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		s.tx = tx
		err := fc(tx)
		s.tx = nil
		return err
	}, opts...)
}
