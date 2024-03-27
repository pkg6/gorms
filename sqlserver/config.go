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

func (c *Config) GetName() string {
	return c.Name
}

func (c *Config) DB() (*gorm.DB, error) {
	return gorm.Open(sqlserver.New(sqlserver.Config{
		DSN:               c.DSN,
		DefaultStringSize: c.DefaultStringSize,
	}), c.Config.GORMConfig())
}
