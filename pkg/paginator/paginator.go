// Package paginator provides a simple paginator implementation for gorm. It
// also supports configuring the paginator via http.Request query params.
package paginator

import (
	"gorm.io/gorm"
)

// type PaginatorResult interface {
// 	IsPaginatorResult()
// }

// func (Result[T]) IsPaginatorResult() {}

// func (*Result[T]) IsPaginatorResult() {}

// DefaultLimit defines the default limit for paginated queries. This is a
// variable so that users can configure it at runtime.
var DefaultLimit = 3

// Paginator defines the interface for a paginator.
type Paginator[T any] interface {
	// Paginate takes a value as arguments and returns a paginated result
	// containing records of the value type.
	Paginate(interface{}) (*Result[T], error)
}

type Params struct {
	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`
}

// paginator defines a paginator.
type paginator[T any] struct {
	db     *gorm.DB
	limit  int
	page   int
	offset int
	order  []string
}

// countResult defines the result of the count query executed by the paginator.
type countResult struct {
	total int64
	err   error
}

// Result defines a paginated result.
type Result[T any] struct {
	CurrentPage    int         `json:"currentPage" gqlgen:"currentPage"`
	Offset         int         `json:"offset" gqlgen:"offset"`
	MaxPage        int         `json:"maxPage" gqlgen:"maxPage"`
	RecordsPerPage int         `json:"recordsPerPage" gqlgen:"recordsPerPage"`
	TotalRecords   int64       `json:"totalRecords" gqlgen:"totalRecords"`
	Records        T           `json:"records" gqlgen:"records"`
	Meta           interface{} `json:"meta" gqlgen:"meta"`
}

// New create a new value of the Paginator type. It expects a gorm DB handle
// and pagination options.
//
//	var v []SomeModel
//	p := gorm-paginator.New(db, gorm-paginator.WithPage(2))
//	res, err := p.Paginate(&v)
func New[T any](db *gorm.DB, options ...Option) Paginator[T] {
	tx := db.Session(&gorm.Session{})
	p := &paginator[any]{
		db:     tx,
		page:   1,
		offset: 0,
		limit:  DefaultLimit,
		order:  make([]string, 0),
	}

	for _, option := range options {
		option(p)
	}

	pp := &paginator[T]{
		db:     tx,
		page:   p.page,
		offset: p.offset,
		limit:  p.limit,
		order:  p.order,
	}

	return pp
}

// Paginate is a convenience wrapper for the paginator.
//
//	var v []SomeModel
//	res, err := gorm-paginator.Paginate(db, &v, gorm-paginator.WithPage(2))
func Paginate[T any](
	db *gorm.DB,
	value interface{},
	options ...Option,
) (*Result[T], error) {
	return New[T](db, options...).Paginate(value)
}

// Paginate implements the Paginator interface.
func (p *paginator[T]) Paginate(value interface{}) (*Result[T], error) {
	db := p.prepareDB()

	c := make(chan countResult, 1)

	go countRecords(db, value, c)

	err := db.Limit(p.limit).Offset(p.getOffset()).Find(value).Error

	if err != nil {
		<-c
		return nil, err
	}

	return p.result(value, <-c)
}

// prepareDB prepares the statement by adding the order clauses.
func (p *paginator[T]) prepareDB() *gorm.DB {
	db := p.db

	for _, o := range p.order {
		db = db.Order(o)
	}

	return db
}

// offset computes the offset used for the paginated query.
func (p *paginator[T]) getOffset() int {
	if p.offset > 0 {
		return p.offset
	}

	return (p.page - 1) * p.limit
}

// countRecords counts the result rows for given query and returns the result
// in the provided channel.
func countRecords(db *gorm.DB, value interface{}, c chan<- countResult) {
	var result countResult
	// TODO: new session done in constructor
	// Session(&gorm.Session{NewDB: true}).
	result.err = db.Model(value).Distinct("id").Count(&result.total).Error
	c <- result
}

// result creates a new Result out of the retrieved value and the count query
func (p *paginator[T]) result(value interface{}, c countResult) (*Result[T], error) {
	if c.err != nil {
		return nil, c.err
	}

	maxPageF := float64(c.total) / float64(p.limit)
	maxPage := int(maxPageF)

	if float64(maxPage) < maxPageF {
		maxPage++
	} else if maxPage == 0 {
		maxPage = 1
	}

	return &Result[T]{
		TotalRecords:   c.total,
		Records:        value.(T),
		Offset:         p.offset,
		CurrentPage:    p.page,
		RecordsPerPage: p.limit,
		MaxPage:        maxPage,
	}, nil
}

// IsLastPage returns true if the current page of the result is the last page.
func (r *Result[T]) IsLastPage() bool {
	return r.CurrentPage >= r.MaxPage
}

// IsFirstPage returns true if the current page of the result is the first page.
func (r *Result[T]) IsFirstPage() bool {
	return r.CurrentPage <= 1
}
