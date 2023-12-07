package scopes

import (
	"gorm.io/gorm"
)

// ScopesPage
//
//db.Scopes(Paginate(r)).Find(&users)
func ScopesPage(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page >= 0 {
			offset := (page - 1) * size
			return db.Offset(offset).Limit(size)
		}
		return db
	}
}
