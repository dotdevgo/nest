package crud

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type (
	Filter struct {
		Name     string
		Operator string
		Value    string
		Expr     string
	}
)

// Apply godoc
func (c *Filter) Apply(db *gorm.DB) (*gorm.DB, error) {
	// TODO: better input validation
	if c.Value == "" || c.Operator == "" {
		return db, nil
	}

	switch c.Operator {
	case "$cont":
		sql := fmt.Sprintf("%s = ?", c.Name)
		db = db.Where(sql, c.Value)
	case "$necont":
		sql := fmt.Sprintf("%s LIKE ?", c.Name)
		db = db.Where(sql, "%"+c.Value+"%")
	default:
		return db, errors.New("invalid filter: " + c.Name)
	}

	return db, nil
}

// Exclude fields from request query params
// var excludeNames = []string{"page", "limit", "order"}

// convertQueryParamToCriteria godoc
//func convertQueryParamToCriteria(name string, data string) CriteriaOption {
//	var operator = "$eq"
//	var parts = strings.Split(data, "||")
//
//	// Operator present in value. Ex: $eq||some value
//	if len(parts) == 2 {
//		operator = parts[0]
//		data = parts[1]
//	}
//
//	return CriteriaOption{
//		Field:    name,
//		Operator: operator,
//		Value:    data,
//	}
//}

// WithCriteria godoc
//func WithCriteria(criteria CriteriaList) Option {
//	return func(db *gorm.DB) *gorm.DB {
//		for _, item := range criteria {
//			db = item.Apply(db)
//			//db.Where(item.SqlExpr())
//		}
//		return db
//	}
//}

// WithRequest godoc
//func WithRequest(r *http.Request) Option {
//	var criteria = Criteria{}
//
//	for name, val := range r.URL.Query() {
//		if slice.ContainsString(excludeNames, name) {
//			continue
//		}
//
//		criteria = append(
//			criteria,
//			convertQueryParamToCriteria(name, strings.Join(val, "")),
//		)
//	}
//
//	return WithCriteria(criteria)
//}

//func NewCriteria(data map[string]interface{}) Criteria {
//	var criteria = Criteria{}
//
//	for name, val := range data {
//		criteria = append(criteria, CriteriaOption{
//			name,
//			"$eq",
//			val,
//			"",
//		})
//	}
//
//	return criteria
//}
//func getSqlOperator(operator string, data string) (string, string) {
//	switch operator {
//	case "$cont":
//		return "LIKE", "%" + data + "%"
//	}
//
//	return "=", data
//}

//var sqlOperator = "="
//var sqlValue = val
//var parts = strings.Split(val.(string), "||")
//
//if len(parts) == 2 {
//	sqlValue = parts[1]
//
//	switch parts[0] {
//	case "$eq":
//		sqlOperator = "="
//		break
//	case "$cont":
//		sqlOperator = "LIKE"
//		sqlValue = "%" + parts[1] + "%"
//		break
//	}
//}
//db = db.Where(name+" "+sqlOperator+" ?", sqlValue)
