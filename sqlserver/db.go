package sqlserver

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// OpenDSN
//sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm
func OpenDSN(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(dsn), new(gorms.GORMConfig).GORMConfig())
}
func Open(sqlserverConfig sqlserver.Config, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(sqlserver.New(sqlserverConfig), config)
}
