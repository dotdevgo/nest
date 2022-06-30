package crud

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/psampaz/slice"
	"gorm.io/gorm"
)

// Exclude fields from request query params
var excludeNames = []string{"page", "limit", "order"}

type (
	CriteriaOption struct {
		Field    string
		Operator string
		Value    interface{}
		Expr     string
	}
	CriteriaMap map[string]CriteriaOption
	Criteria    []CriteriaOption
)

// ScopeOrderBy godoc
func ScopeOrderBy(column string, order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", column, strings.ToUpper(order)))
	}
}

// ScopeById godoc
func ScopeById(result interface{}, id interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		table := GetTableName(db, result)
		str, ok := id.(string)
		if ok {
			_, err := uuid.Parse(str)
			if err == nil {
				return db.Where(table+".id = ?", id)
			}
		}
		return db.Where(table+".pk = ?", id)
	}
}

// WithCriteria godoc
func WithCriteria(criteria Criteria) Option {
	return func(db *gorm.DB) *gorm.DB {
		for _, item := range criteria {
			db = item.Apply(db)
			//db.Where(item.SqlExpr())
		}
		return db
	}
}

// WithRequest godoc
func WithRequest(r *http.Request) Option {
	var criteria = Criteria{}

	for name, val := range r.URL.Query() {
		if slice.ContainsString(excludeNames, name) {
			continue
		}

		criteria = append(
			criteria,
			convertQueryParamToCriteria(name, strings.Join(val, "")),
		)
	}

	return WithCriteria(criteria)
}

// Apply godoc
func (c *CriteriaOption) Apply(db *gorm.DB) *gorm.DB {
	var sql = c.Expr
	var value = c.Value.(string)
	var operator = "="
	var isNot = false

	if value == "" || c.Operator == "" {
		return db
	}

	switch c.Operator {
	case "$cont":
		operator = "LIKE"
		value = "%" + value + "%"
	case "$necont":
		operator = "LIKE"
		value = "%" + value + "%"
		isNot = true
	}

	if sql == "" {
		sql = c.Field + " " + operator + " ?"
	}

	if isNot {
		return db.Not(sql, value)
	}

	return db.Where(sql, value)
}

// convertQueryParamToCriteria godoc
func convertQueryParamToCriteria(name string, data string) CriteriaOption {
	var operator = "$eq"
	var parts = strings.Split(data, "||")

	// Operator present in value. Ex: $eq||some value
	if len(parts) == 2 {
		operator = parts[0]
		data = parts[1]
	}

	return CriteriaOption{
		Field:    name,
		Operator: operator,
		Value:    data,
	}
}

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
