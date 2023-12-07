package scopes

import "gorm.io/gorm"

type IScopes interface {
	Scopes() func(db *gorm.DB) *gorm.DB
}
