package sqlserver

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Config struct {
	Name              string
	DSN               string
	DefaultStringSize int
	Config            *gorms.GORMConfig
}

func (m *Config) GetName() string {
	return m.Name
}

func (m *Config) DB() (*gorm.DB, error) {
	return gorm.Open(sqlserver.New(sqlserver.Config{
		DSN:               m.DSN,
		DefaultStringSize: m.DefaultStringSize,
	}), m.Config.GORMConfig())
}
