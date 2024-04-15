package sqlite

import (
	"github.com/pkg6/gorms"
	"gorm.io/gorm"
)

type Config struct {
	Name   string
	DSN    string
	Config *gorms.GORMConfig
}

func (c *Config) GetName() string {
	return c.Name
}

func (c *Config) DB() (*gorm.DB, error) {
	return Open(c.DSN, c.Config.GORMConfig())
}
