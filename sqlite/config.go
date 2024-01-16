package sqlite

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Name   string
	DSN    string
	Config *gorms.GORMConfig
}

func (m *Config) GetName() string {
	return m.Name
}

func (m *Config) DB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(m.DSN), m.Config.GORMConfig())
}
