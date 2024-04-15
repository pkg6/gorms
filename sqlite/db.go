package sqlite

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDSN(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), new(gorms.GORMConfig).GORMConfig())
}
func Open(dsn string, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), config)
}
