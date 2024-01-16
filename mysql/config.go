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

func (m *Config) GetName() string {
	return m.Name
}
func (m *Config) DB() (*gorm.DB, error) {
	dsnConf, _ := mysqldriver.ParseDSN(m.DSN)
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                           m.DSN,
		DSNConfig:                     dsnConf,
		SkipInitializeWithVersion:     m.SkipInitializeWithVersion,
		DefaultStringSize:             m.DefaultStringSize,
		DefaultDatetimePrecision:      m.DefaultDatetimePrecision,
		DisableWithReturning:          m.DisableWithReturning,
		DisableDatetimePrecision:      m.DisableDatetimePrecision,
		DontSupportRenameIndex:        m.DontSupportRenameIndex,
		DontSupportRenameColumn:       m.DontSupportRenameColumn,
		DontSupportForShareClause:     m.DontSupportForShareClause,
		DontSupportNullAsDefaultValue: m.DontSupportNullAsDefaultValue,
		DontSupportRenameColumnUnique: m.DontSupportRenameColumnUnique,
	}), m.Config.GORMConfig())
}
