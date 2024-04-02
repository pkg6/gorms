package postgres

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// OpenDSN
//host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
func OpenDSN(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), new(gorms.GORMConfig).GORMConfig())
}

func Open(postgresConfig postgres.Config, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgresConfig), config)
}
