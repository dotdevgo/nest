package crud

import (
	// "dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/paginator"
	"net/http"
)

type (
	PaginatorMeta interface {
		// IsPaginatorMeta()
	}
	PaginatorCursor struct {
		Offset *int    `json:"offset"` // NOT USED
		Page   *int    `json:"page"`
		Limit  *int    `json:"limit"`
		Order  *string `json:"order"`
	}
)

// Paginate godoc
func (s *Crud[T]) Paginate(result interface{}, pagination []paginator.Option, options ...Option) (*paginator.Result[T], error) {
	var stmt = s.Stmt(options...)

	return paginator.Paginate[T](stmt, result, pagination...)
}

// WithNest godoc
func WithNest(ctx nest.Context) []paginator.Option {
	return []paginator.Option{
		paginator.WithRequest(ctx.Request()),
	}
}

// WithHttpRequest godoc
func WithHttpRequest(req *http.Request) []paginator.Option {
	return []paginator.Option{
		paginator.WithRequest(req),
	}
}

// WithCursor godoc
func WithCursor(cursor *PaginatorCursor) []paginator.Option {
	var pgr []paginator.Option

	if cursor.Page != nil {
		pgr = append(pgr, paginator.WithPage(*cursor.Page))
	}
	if cursor.Limit != nil {
		pgr = append(pgr, paginator.WithLimit(*cursor.Limit))
	}
	if cursor.Order != nil {
		pgr = append(pgr, paginator.WithOrder(*cursor.Order))
	}
	//else {
	//	pgr = append(pgr, paginator.WithOrder(entity.TableFormTemplate+".created_at desc"))
	//}

	return pgr
}

//PaginatorResult struct {
//	CurrentPage    int           `json:"currentPage"`
//	MaxPage        int           `json:"maxPage"`
//	RecordsPerPage int           `json:"recordsPerPage"`
//	TotalRecords   int64         `json:"totalRecords"`
//	Meta           PaginatorMeta `json:"meta"`
//	Records        interface{}   `json:"records"`
//}
// NewPaginatorCursor godoc
//func NewPaginatorCursor(pg *paginator.Result) *PaginatorResult {
//	return &PaginatorResult{
//		CurrentPage:    pg.CurrentPage,
//		MaxPage:        pg.MaxPage,
//		RecordsPerPage: pg.RecordsPerPage,
//		TotalRecords:   pg.TotalRecords,
//	}
//}
