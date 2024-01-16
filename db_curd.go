package gorms

import (
	"gorm.io/gorm"
)

// Inc
//字段加 5
// Inc(db.Table("test").Where("id = ?",1), "limit", 5)
func Inc(db *gorm.DB, column string, values ...any) error {
	var value any = 1
	if len(values) > 0 {
		value = values[0]
	}
	return db.UpdateColumn(column, gorm.Expr("`"+column+"` + ?", value)).Error
}

// Dec
//字段减 1
//Dec(db.Table("test").Where("id = ?",1), "limit", 5)
func Dec(db *gorm.DB, column string, values ...any) error {
	var value any = 1
	if len(values) > 0 {
		value = values[0]
	}
	return db.UpdateColumn(column, gorm.Expr("`"+column+"` - ?", value)).Error
}
