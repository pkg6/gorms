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

func (c *Config) GetName() string {
	return c.Name
}

func (c *Config) DB() (*gorm.DB, error) {
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  c.DSN,
		WithoutQuotingCheck:  c.WithoutQuotingCheck,
		PreferSimpleProtocol: c.PreferSimpleProtocol,
		WithoutReturning:     c.WithoutReturning,
	}), c.Config.GORMConfig())
}
