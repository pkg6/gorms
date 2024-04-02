package mysql

import (
	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/pkg6/gorms"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Name                          string
	DSN                           string
	SkipInitializeWithVersion     bool
	DefaultStringSize             uint
	DefaultDatetimePrecision      *int
	DisableWithReturning          bool
	DisableDatetimePrecision      bool
	DontSupportRenameIndex        bool
	DontSupportRenameColumn       bool
	DontSupportForShareClause     bool
	DontSupportNullAsDefaultValue bool
	DontSupportRenameColumnUnique bool
	Config                        *gorms.GORMConfig
}

func (c *Config) GetName() string {
	return c.Name
}
func (c *Config) DB() (*gorm.DB, error) {
	dsnConf, _ := mysqldriver.ParseDSN(c.DSN)
	return Open(mysql.Config{
		DSN:                           c.DSN,
		DSNConfig:                     dsnConf,
		SkipInitializeWithVersion:     c.SkipInitializeWithVersion,
		DefaultStringSize:             c.DefaultStringSize,
		DefaultDatetimePrecision:      c.DefaultDatetimePrecision,
		DisableWithReturning:          c.DisableWithReturning,
		DisableDatetimePrecision:      c.DisableDatetimePrecision,
		DontSupportRenameIndex:        c.DontSupportRenameIndex,
		DontSupportRenameColumn:       c.DontSupportRenameColumn,
		DontSupportForShareClause:     c.DontSupportForShareClause,
		DontSupportNullAsDefaultValue: c.DontSupportNullAsDefaultValue,
		DontSupportRenameColumnUnique: c.DontSupportRenameColumnUnique,
	}, c.Config.GORMConfig())
}
