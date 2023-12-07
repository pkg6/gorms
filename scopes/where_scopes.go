package scopes

import (
	"github.com/pkg6/gorms"
	"gorm.io/gorm"
	"strings"
)

type Criteria struct {
	Column    string
	Operator  string
	Value     []any
	Connector string
}

func (c *Criteria) SQLArgs() (string, []any) {
	var query string
	switch c.Operator {
	case gorms.OperatorBetween, gorms.OperatorNotBetween:
		query = c.Column + " " + c.Operator + " ? AND ?"
	case gorms.OperatorIsNull, gorms.OperatorIsNotNull:
		query = c.Column + " " + c.Operator
	default:
		query = c.Column + " " + c.Operator + " ? "
	}
	return query, c.Value
}

type WhereScopes struct {
	Criterias []*Criteria
}

func NewWhere() *WhereScopes {
	return &WhereScopes{
		Criterias: []*Criteria{},
	}
}

// Eq 等于 =
func (q *WhereScopes) Eq(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorEq, []any{value}, gorms.ConnectorAnd)
	return q
}

// Ne 不等于 !=
func (q *WhereScopes) Ne(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorNe, []any{value}, gorms.ConnectorAnd)
	return q
}

// Gt 大于 >
func (q *WhereScopes) Gt(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorGt, []any{value}, gorms.ConnectorAnd)
	return q
}

// Ge 大于等于 >=
func (q *WhereScopes) Ge(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorGe, []any{value}, gorms.ConnectorAnd)
	return q
}

// Lt 小于 <
func (q *WhereScopes) Lt(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorLt, []any{value}, gorms.ConnectorAnd)
	return q
}

// Le 小于等于 <=
func (q *WhereScopes) Le(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorLe, []any{value}, gorms.ConnectorAnd)
	return q
}

// Like 模糊 LIKE '%值%'
func (q *WhereScopes) Like(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorLike, []any{value}, gorms.ConnectorAnd)
	return q
}

// NotLike 非模糊 NOT LIKE '%值%'
func (q *WhereScopes) NotLike(column string, value any) *WhereScopes {
	q.Where(column, gorms.OperatorNotLike, []any{value}, gorms.ConnectorAnd)
	return q
}

//IsNull 是否为空 字段 IS NULL
func (q *WhereScopes) IsNull(column string) *WhereScopes {
	q.Where(column, gorms.OperatorIsNull, []any{""}, gorms.ConnectorAnd)
	return q
}

// IsNotNull 是否非空 字段 IS NOT NULL
func (q *WhereScopes) IsNotNull(column string) *WhereScopes {
	q.Where(column, gorms.OperatorIsNotNull, []any{""}, gorms.ConnectorAnd)
	return q
}

// In 字段 IN (值1, 值2, ...)
func (q *WhereScopes) In(column string, value []any) *WhereScopes {
	q.Where(column, gorms.OperatorIn, value, gorms.ConnectorAnd)
	return q
}

// NotIn 字段 NOT IN (值1, 值2, ...)
func (q *WhereScopes) NotIn(column string, value []any) *WhereScopes {
	q.Where(column, gorms.OperatorNotIn, value, gorms.ConnectorAnd)
	return q
}

// Between BETWEEN 值1 AND 值2
func (q *WhereScopes) Between(column string, s any, e any) *WhereScopes {
	q.Where(column, gorms.OperatorBetween, []any{s, e}, gorms.ConnectorAnd)
	return q
}

// NotBetween NOT BETWEEN 值1 AND 值2
func (q *WhereScopes) NotBetween(column string, s any, e any) *WhereScopes {
	q.Where(column, gorms.OperatorNotBetween, []any{s, e}, gorms.ConnectorAnd)
	return q
}

func (q *WhereScopes) OrWhere(column string, operator string, value []any) *WhereScopes {
	q.Where(column, operator, value, gorms.ConnectorOr)
	return q
}

func (q *WhereScopes) Where(column string, operator string, value []any, connector string) *WhereScopes {
	q.Criterias = append(q.Criterias, &Criteria{Column: column, Operator: operator, Value: value, Connector: connector})
	return q
}

func (q *WhereScopes) SQLArgs() (string, []any) {
	var (
		sqla  []string
		argss []any
	)
	for _, criteria := range q.Criterias {
		sql, args := criteria.SQLArgs()
		if len(sqla) <= 0 {
			sqla = append(sqla, sql+" ")
		} else {
			sqla = append(sqla, " "+criteria.Connector+" "+sql+" ")
		}
		argss = append(argss, args...)
	}
	return strings.Join(sqla, ""), argss
}

func (q *WhereScopes) Scopes() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		sql, args := q.SQLArgs()
		if sql != "" {
			return db.Where(sql, args...)
		}
		return db
	}
}
