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

func (m *Config) GetName() string {
	return m.Name
}

func (m *Config) DB() (*gorm.DB, error) {
	return gorm.Open(clickhouse.New(clickhouse.Config{
		DSN:                          m.DSN,
		DisableDatetimePrecision:     m.DisableDatetimePrecision,
		DontSupportRenameColumn:      m.DontSupportRenameColumn,
		DontSupportColumnPrecision:   m.DontSupportColumnPrecision,
		DontSupportEmptyDefaultValue: m.DontSupportEmptyDefaultValue,
		SkipInitializeWithVersion:    m.SkipInitializeWithVersion,
		DefaultGranularity:           m.DefaultGranularity,
		DefaultCompression:           m.DefaultCompression,
		DefaultIndexType:             m.DefaultIndexType,
		DefaultTableEngineOpts:       m.DefaultTableEngineOpts,
	}), m.Config.GORMConfig())
}
