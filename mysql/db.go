package mysql

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// OpenDSN
//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func OpenDSN(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), new(gorms.GORMConfig).GORMConfig())
}

func Open(mysqlConfig mysql.Config, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysqlConfig), config)
}
