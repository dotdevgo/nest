package paginator

import (
	"dotdev/nest"
	"net/http"
	"strconv"
	"strings"
)

// Option defines the signature of a paginator option function.
type Option func(p *paginator[any])

// ParamNames defines a type to configure names of query parameters to use from
// a http.Request. If a field is set to the empty string, it will not be used.
type ParamNames struct {
	Page   string
	Limit  string
	Order  string
	Offset string
}

// DefaultParamNames specifies the query parameter names to use from
// a http.Request by default when the WithRequest option is uses. This can be
// overridden at runtime.
var DefaultParamNames = ParamNames{"page", "limit", "order", "offset"}

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

// WithHttpRequest godoc
func WithHttpRequest(req *http.Request) []Option {
	return []Option{
		WithRequest(req),
	}
}

// WithCursor godoc
func WithCursor(cursor *PaginatorCursor) []Option {
	var pgr []Option

	if cursor.Page != nil {
		pgr = append(pgr, WithPage(*cursor.Page))
	}
	if cursor.Limit != nil {
		pgr = append(pgr, WithLimit(*cursor.Limit))
	}
	if cursor.Order != nil {
		pgr = append(pgr, WithOrder(*cursor.Order))
	}
	//else {
	//	pgr = append(pgr, WithOrder(entity.TableFormTemplate+".created_at desc"))
	//}

	return pgr
}

// WithContext godoc
func WithContext(ctx nest.Context) []Option {
	return []Option{
		WithRequest(ctx.Request()),
	}
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
//func NewPaginatorCursor(pg *Result) *PaginatorResult {
//	return &PaginatorResult{
//		CurrentPage:    pg.CurrentPage,
//		MaxPage:        pg.MaxPage,
//		RecordsPerPage: pg.RecordsPerPage,
//		TotalRecords:   pg.TotalRecords,
//	}
//}

// WithOffset configures the offset of the
//
//	gorm-Paginate(db, &v, gorm-WithOffset(2))
func WithOffset(offset int) Option {
	return func(p *paginator[any]) {
		if offset > 0 {
			p.offset = offset
		}
	}
}

// WithPage configures the page of the
//
//	gorm-Paginate(db, &v, gorm-WithPage(2))
func WithPage(page int) Option {
	return func(p *paginator[any]) {
		if page > 0 {
			p.page = page
		}
	}
}

// WithLimit configures the limit of the
//
//	gorm-Paginate(db, &v, gorm-WithLimit(10))
func WithLimit(limit int) Option {
	return func(p *paginator[any]) {
		if limit > 0 {
			p.limit = limit
		}
	}
}

// WithOrder configures the order of the
//
//	gorm-Paginate(db, &v, gorm-WithOrder("name DESC", "id"))
func WithOrder(order ...string) Option {
	return func(p *paginator[any]) {
		p.order = filterNonEmpty(order)
	}
}

// WithRequest configures the paginator from a *http.Request.
//
//	gorm-Paginate(db, &v, gorm-WithRequest(request))
//
//	gorm-Paginate(db, &v, gorm-WithRequest(request, gorm-ParamNames{
//	    Page: "page",
//	    Limit: "",       // Disable limit query param.
//	    Order: "order",
//	}))
func WithRequest(r *http.Request, paramNames ...ParamNames) Option {
	params := DefaultParamNames

	if len(paramNames) > 0 {
		params = paramNames[0]
	}

	return func(p *paginator[any]) {
		if value, ok := getQueryParam(r, params.Offset); ok {
			if offset, err := strconv.Atoi(value); err == nil {
				WithOffset(offset)(p)
			}
		}

		if value, ok := getQueryParam(r, params.Page); ok {
			if page, err := strconv.Atoi(value); err == nil {
				WithPage(page)(p)
			}
		}

		if value, ok := getQueryParam(r, params.Limit); ok {
			if limit, err := strconv.Atoi(value); err == nil {
				WithLimit(limit)(p)
			}
		}

		if value, ok := getQueryParam(r, params.Order); ok {
			if order := strings.TrimSpace(value); len(order) > 0 {
				WithOrder(strings.Split(order, ",")...)(p)
			}
		}
	}
}

// WithRequest configures the paginator from a *http.Request.
//
//	gorm-Paginate(db, &v, gorm-WithRequest(request))
//
//	gorm-Paginate(db, &v, gorm-WithRequest(request, gorm-ParamNames{
//	    Page: "page",
//	    Limit: "",       // Disable limit query param.
//	    Order: "order",
//	}))
func NewParams(r *http.Request, paramNames ...ParamNames) *Params {
	params := DefaultParamNames

	p := &Params{}

	if len(paramNames) > 0 {
		params = paramNames[0]
	}

	if value, ok := getQueryParam(r, params.Offset); ok {
		if offset, err := strconv.Atoi(value); err == nil {
			// WithOffset(offset)(p)
			p.Offset = &offset
		}
	}

	// if value, ok := getQueryParam(r, params.Page); ok {
	// 	if page, err := strconv.Atoi(value); err == nil {
	// 		WithPage(page)(p)
	// 	}
	// }

	if value, ok := getQueryParam(r, params.Limit); ok {
		if limit, err := strconv.Atoi(value); err == nil {
			// WithLimit(limit)(p)
			p.Limit = &limit
		}
	}

	// if value, ok := getQueryParam(r, params.Order); ok {
	// 	if order := strings.TrimSpace(value); len(order) > 0 {
	// 		WithOrder(strings.Split(order, ",")...)(p)
	// 	}
	// }

	return p
}

// getQueryParam gets the first query param matching key from the request.
// Returns empty string of key or param value is empty. Second return value
// indicates wether the param was present in the query or not.
func getQueryParam(r *http.Request, key string) (string, bool) {
	if key == "" {
		return "", false
	}

	if values, ok := r.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}

	return "", false
}

// filterNonEmpty filters out all elements that are either empty or contain
// solely whitespace characters.
func filterNonEmpty(elements []string) []string {
	nonEmpty := make([]string, 0)

	for _, el := range elements {
		if el = strings.TrimSpace(el); len(el) > 0 {
			nonEmpty = append(nonEmpty, el)
		}
	}

	return nonEmpty
}
