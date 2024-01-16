package gorms

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
