package clickhouse

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// OpenDSN
//clickhouse://gorm:gorm@localhost:9942/gorm?dial_timeout=10s&read_timeout=20s
func OpenDSN(dsn string) (*gorm.DB, error) {
	return gorm.Open(clickhouse.Open(dsn), new(gorms.GORMConfig).GORMConfig())
}

func Open(clickhouseConfig clickhouse.Config, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(clickhouse.New(clickhouseConfig), config)
}
