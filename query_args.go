package gorms

import (
	"gorm.io/gorm"
	"strings"
)

const (
	ConnectorAnd = "AND"
	ConnectorOr  = "OR"
)

const (
	OperatorIn         = "IN"
	OperatorNot        = "NOT"
	OperatorLike       = "LIKE"
	OperatorEq         = "="
	OperatorNe         = "<>"
	OperatorGt         = ">"
	OperatorGe         = ">="
	OperatorLt         = "<"
	OperatorLe         = "<="
	OperatorIsNull     = "IS NULL"
	OperatorIsNotNull  = "IS NOT NULL"
	OperatorBetween    = "BETWEEN"
	OperatorNotBetween = OperatorNot + " " + OperatorBetween
	OperatorNotIn      = OperatorNot + " " + OperatorIn
	OperatorNotLike    = OperatorNot + " " + OperatorLike
)

const (
	OrderByDesc = "DESC"
	OrderByAsc  = "ASC"
)

const (
	SUM   = "SUM"
	AVG   = "AVG"
	MAX   = "MAX"
	MIN   = "MIN"
	COUNT = "COUNT"
)

// ScopesWhereQueryArgs
//args = NewQueryArgs()
//db.Scopes(ScopesWhereQueryArgs(args)).Find(&users)
func ScopesWhereQueryArgs(queryArgs *QueryArgs) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query, arg := queryArgs.WhereQueryArgs()
		if query != "" {
			return db.Where(query, arg...)
		}
		return db
	}
}

func ScopesCallbackWhereQueryArgs(fns ...func(queryArgs *QueryArgs)) func(db *gorm.DB) *gorm.DB {
	args := NewQueryArgs()
	for _, fn := range fns {
		fn(args)
	}
	return ScopesWhereQueryArgs(args)
}

func NewQueryArgs() *QueryArgs {
	return &QueryArgs{
		criteria: []*Criteria{},
	}
}

type IQueryArgs interface {
	QueryArgs() (string, []any)
}

type QueryArgs struct {
	//条件筛选
	criteria []*Criteria
}

// Eq 等于 =
func (q *QueryArgs) Eq(column string, value any) *QueryArgs {
	q.Where(column, OperatorEq, []any{value}, ConnectorAnd)
	return q
}

// Ne 不等于 !=
func (q *QueryArgs) Ne(column string, value any) *QueryArgs {
	q.Where(column, OperatorNe, []any{value}, ConnectorAnd)
	return q
}

// Gt 大于 >
func (q *QueryArgs) Gt(column string, value any) *QueryArgs {
	q.Where(column, OperatorGt, []any{value}, ConnectorAnd)
	return q
}

// Ge 大于等于 >=
func (q *QueryArgs) Ge(column string, value any) *QueryArgs {
	q.Where(column, OperatorGe, []any{value}, ConnectorAnd)
	return q
}

// Lt 小于 <
func (q *QueryArgs) Lt(column string, value any) *QueryArgs {
	q.Where(column, OperatorLt, []any{value}, ConnectorAnd)
	return q
}

// Le 小于等于 <=
func (q *QueryArgs) Le(column string, value any) *QueryArgs {
	q.Where(column, OperatorLe, []any{value}, ConnectorAnd)
	return q
}

// Like 模糊 LIKE '%值%'
func (q *QueryArgs) Like(column string, value any) *QueryArgs {
	q.Where(column, OperatorLike, []any{value}, ConnectorAnd)
	return q
}

// NotLike 非模糊 NOT LIKE '%值%'
func (q *QueryArgs) NotLike(column string, value any) *QueryArgs {
	q.Where(column, OperatorNotLike, []any{value}, ConnectorAnd)
	return q
}

//IsNull 是否为空 字段 IS NULL
func (q *QueryArgs) IsNull(column string) *QueryArgs {
	q.Where(column, OperatorIsNull, []any{""}, ConnectorAnd)
	return q
}

// IsNotNull 是否非空 字段 IS NOT NULL
func (q *QueryArgs) IsNotNull(column string) *QueryArgs {
	q.Where(column, OperatorIsNotNull, []any{""}, ConnectorAnd)
	return q
}

// In 字段 IN (值1, 值2, ...)
func (q *QueryArgs) In(column string, value []any) *QueryArgs {
	q.Where(column, OperatorIn, value, ConnectorAnd)
	return q
}

// NotIn 字段 NOT IN (值1, 值2, ...)
func (q *QueryArgs) NotIn(column string, value []any) *QueryArgs {
	q.Where(column, OperatorNotIn, value, ConnectorAnd)
	return q
}

// Between BETWEEN 值1 AND 值2
func (q *QueryArgs) Between(column string, s any, e any) *QueryArgs {
	q.Where(column, OperatorBetween, []any{s, e}, ConnectorAnd)
	return q
}

// NotBetween NOT BETWEEN 值1 AND 值2
func (q *QueryArgs) NotBetween(column string, s any, e any) *QueryArgs {
	q.Where(column, OperatorNotBetween, []any{s, e}, ConnectorAnd)
	return q
}

func (q *QueryArgs) OrWhere(column string, operator string, value []any) *QueryArgs {
	q.Where(column, operator, value, ConnectorOr)
	return q
}

func (q *QueryArgs) Where(column string, operator string, value []any, connector string) *QueryArgs {
	q.criteria = append(q.criteria, &Criteria{Column: column, Operator: operator, Value: value, Connector: connector})
	return q
}

// WhereQueryArgs  where 条件筛选
//args = NewSQLArgs()
//db.Scopes(ScopesWhereQueryArgs(args)).Find(&users)
func (q *QueryArgs) WhereQueryArgs() (string, []any) {
	var (
		sqls []string
		args []any
	)
	for _, criteria := range q.criteria {
		sql, arg := criteria.QueryArgs()
		if len(sqls) <= 0 {
			sqls = append(sqls, sql+" ")
		} else {
			sqls = append(sqls, " "+criteria.Connector+" "+sql+" ")
		}
		args = append(args, arg...)
	}
	return strings.Join(sqls, ""), args
}

type Criteria struct {
	Column    string
	Operator  string
	Value     []any
	Connector string
}

func (c *Criteria) QueryArgs() (string, []any) {
	var query string
	switch c.Operator {
	case OperatorBetween, OperatorNotBetween:
		query = c.Column + " " + c.Operator + " ? AND ?"
	case OperatorIsNull, OperatorIsNotNull:
		query = c.Column + " " + c.Operator
	case OperatorIn, OperatorNotIn:
		buf := new(strings.Builder)
		for range c.Value {
			buf.WriteString(" ?,")
		}
		var char string
		if buf.Len() != 0 {
			str := buf.String()
			char = str[:len(str)-1]
		}
		query = c.Column + " " + c.Operator + " ( " + char + " )"
	default:
		query = c.Column + " " + c.Operator + " ? "
	}
	return query, c.Value
}
