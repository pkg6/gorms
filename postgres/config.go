package postgres

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Name                 string
	DSN                  string
	WithoutQuotingCheck  bool
	PreferSimpleProtocol bool
	WithoutReturning     bool
	Config               *gorms.GORMConfig
}

func (m *Config) GetName() string {
	return m.Name
}

func (m *Config) DB() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  m.DSN,
		WithoutQuotingCheck:  m.WithoutQuotingCheck,
		PreferSimpleProtocol: m.PreferSimpleProtocol,
		WithoutReturning:     m.WithoutReturning,
	}), m.Config.GORMConfig())
}
