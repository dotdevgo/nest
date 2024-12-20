package crud

import (
	"dotdev/orm"

	"github.com/defval/di"
	"gorm.io/gorm"
)

// Repository godoc
type Repository[T orm.Model] struct {
	di.Inject

	*gorm.DB
}

// Find godoc
func (s *Repository[T]) Find(id string) (T, error) {
	var data T
	var uuid = orm.UUIDToBinary(id)

	stmt := s.First(&data, "id = ?", uuid)
	if err := stmt.Error; err != nil {
		return data, err
	}

	return data, nil
}

// FindAll godoc
func (s *Repository[T]) FindAll(result []T, options ...Option) error {
	var stmt = s.CreateQuery(options...)

	return stmt.Find(result).Error
}

// CreateQuery godoc
func (r *Repository[T]) CreateQuery(options ...Option) *gorm.DB {
	var stmt = r.Session(&gorm.Session{NewDB: true})
	for _, option := range options {
		stmt = option(stmt)
	}

	return stmt
}

// NewRepository godoc
func NewRepository[T orm.Model]() *Repository[T] {
	return &Repository[T]{}
}

// Upsert godoc
// func (r *Repository[T]) Upsert(data T) error {
// 	if len(data.GetID()) > 0 {
// 		return r.Save(data).Error
// 	}

// 	return r.Create(data).Error
// }

// Paginate godoc
// func (s *Repository[T]) Paginate(model T, pagination []paginator.Option, options ...Option) (*paginator.Result[[]T], error) {
// 	var stmt = s.createQuery(options...)

// 	return paginator.Paginate[[]T](stmt, model, pagination...)
// }
