package crud

import (
	"dotdev/nest/pkg/nest"
)

type Crud[T Model] struct {
	// db *gorm.DB
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

// Stmt godoc
// func (s *Crud[T]) NewQuery(options ...Option) *gorm.DB {
// 	var stmt = s.db.Session(&gorm.Session{}) // Tx() //NewDB: true

// 	for _, option := range options {
// 		stmt = option(stmt)
// 	}

// 	return stmt
// }

// NewService godoc
//func NewService[T Model](db *gorm.DB) *Crud[T] {
//	return &Crud[T]{db: db}
//	// .Session(&gorm.Session{NewDB: true})
//}

// Find godoc
// @deprecated
// TODO: move to repository
//func (s *Crud[T]) Find(result T, id interface{}, options ...Option) error {
//	stmt := s.Stmt(options...)
//
//	return stmt.Scopes(ScopeById(result, id)).First(result).Error
//}

// FindAll godoc
// @deprecated
// TODO: move to repository
//func (s *Crud[T]) FindAll(result interface{}, options ...Option) error {
//	var stmt = s.Stmt(options...)
//
//	return stmt.Find(result).Error
//}

// Flush godoc
// func (s *Crud[T]) Flush(data T) error {
// 	if data.GetPk() > 0 || data.GetID() != "" {
// 		return s.Tx().Save(data).Error
// 	}

// 	return s.Tx().Create(data).Error
// }

// Tx Get current transaction
//
// TODO: @internal remove refactor
// func (s *Crud[T]) Tx() *gorm.DB {
// 	// if s.tx != nil {
// 	// 	return s.tx
// 	// }
// 	return s.db
// }

// DB Godoc
// @deprecated
// TODO: remove/refactor
// func (s *Crud[T]) DB() *gorm.DB {
// 	return s.db
// }

// Begin godoc
// func (s *Crud[T]) Begin(opts ...*sql.TxOptions) (*gorm.DB, error) {
// 	s.tx = s.db.Begin(opts...)

// 	defer func() {
// 		if r := recover(); r != nil && s.tx != nil {
// 			s.tx.Rollback()
// 		}
// 	}()

// 	if err := s.tx.Error; err != nil {
// 		return nil, err
// 	}

// 	return s.tx, nil
// }

// Commit godoc
// func (s *Crud[T]) Commit() *gorm.DB {
// 	if s.tx == nil {
// 		return nil
// 	}

// 	tx := s.tx.Commit()

// 	s.tx = nil

// 	return tx
// }

// Rollback godoc
// func (s *Crud[T]) Rollback() {
// 	if s.tx == nil {
// 		return
// 	}

// 	s.tx.Rollback()
// 	s.tx = nil
// }

// Transaction godoc
// func (s *Crud[T]) Transaction(fc func(tx *gorm.DB) (err error), opts ...*sql.TxOptions) error {
// 	return s.db.Transaction(func(tx *gorm.DB) error {
// 		s.tx = tx
// 		err := fc(tx)
// 		s.tx = nil
// 		return err
// 	}, opts...)
// }
