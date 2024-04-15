package clickhouse

import (
	"github.com/pkg6/gorms"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type Config struct {
	Name                         string
	DSN                          string
	DisableDatetimePrecision     bool
	DontSupportRenameColumn      bool
	DontSupportColumnPrecision   bool
	DontSupportEmptyDefaultValue bool
	SkipInitializeWithVersion    bool
	DefaultGranularity           int    // 1 granule = 8192 rows
	DefaultCompression           string // default compression algorithm. LZ4 is lossless
	DefaultIndexType             string // index stores extremes of the expression
	DefaultTableEngineOpts       string
	Config                       *gorms.GORMConfig
}

func (c *Config) GetName() string {
	return c.Name
}

func (c *Config) DB() (*gorm.DB, error) {
	return Open(clickhouse.Config{
		DSN:                          c.DSN,
		DisableDatetimePrecision:     c.DisableDatetimePrecision,
		DontSupportRenameColumn:      c.DontSupportRenameColumn,
		DontSupportColumnPrecision:   c.DontSupportColumnPrecision,
		DontSupportEmptyDefaultValue: c.DontSupportEmptyDefaultValue,
		SkipInitializeWithVersion:    c.SkipInitializeWithVersion,
		DefaultGranularity:           c.DefaultGranularity,
		DefaultCompression:           c.DefaultCompression,
		DefaultIndexType:             c.DefaultIndexType,
		DefaultTableEngineOpts:       c.DefaultTableEngineOpts,
	}, c.Config.GORMConfig())
}
