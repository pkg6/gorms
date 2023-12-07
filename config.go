package gorms

import (
	mysqldriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Config struct {
	Mysql      *ConfigMysql      `gorm:"Mysql"`
	SQLite     *ConfigSQLite     `gorm:"SQLite"`
	Postgres   *ConfigPostgres   `gorm:"Postgres"`
	SQLServer  *ConfigSQLServer  `gorm:"SQLServer"`
	Clickhouse *ConfigClickhouse `gorm:"Clickhouse"`
}

var (
	DefaultLogger     = logger.Default
	DefaultGORMConfig = &gorm.Config{
		QueryFields: true,
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: DefaultLogger,
	}
)

type GORMConfig struct {
	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction bool
	// FullSaveAssociations full save associations
	FullSaveAssociations bool
	// DryRun generate sql without execute
	DryRun bool
	// PrepareStmt executes the given query in cached statement
	PrepareStmt bool
	// DisableAutomaticPing
	DisableAutomaticPing bool
	// DisableForeignKeyConstraintWhenMigrating
	DisableForeignKeyConstraintWhenMigrating bool
	// IgnoreRelationshipsWhenMigrating
	IgnoreRelationshipsWhenMigrating bool
	// DisableNestedTransaction disable nested transaction
	DisableNestedTransaction bool
	// AllowGlobalUpdate allow global update
	AllowGlobalUpdate bool
	// QueryFields executes the SQL query with all fields of the table
	QueryFields bool
	// CreateBatchSize default create batch size
	CreateBatchSize int
	// TranslateError enabling error translation
	TranslateError bool
	//tables, columns naming strategy
	NamingStrategy *NamingStrategy
}

type NamingStrategy struct {
	TablePrefix         string
	SingularTable       bool
	NoLowerCase         bool
	IdentifierMaxLength int
}

func (config *GORMConfig) GORMConfig() *gorm.Config {
	if config != nil {
		return &gorm.Config{
			SkipDefaultTransaction:                   config.SkipDefaultTransaction,
			FullSaveAssociations:                     config.FullSaveAssociations,
			DryRun:                                   config.DryRun,
			PrepareStmt:                              config.PrepareStmt,
			DisableAutomaticPing:                     config.DisableAutomaticPing,
			DisableForeignKeyConstraintWhenMigrating: config.DisableForeignKeyConstraintWhenMigrating,
			IgnoreRelationshipsWhenMigrating:         config.IgnoreRelationshipsWhenMigrating,
			DisableNestedTransaction:                 config.DisableNestedTransaction,
			AllowGlobalUpdate:                        config.AllowGlobalUpdate,
			QueryFields:                              config.QueryFields,
			CreateBatchSize:                          config.CreateBatchSize,
			TranslateError:                           config.TranslateError,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:         config.NamingStrategy.TablePrefix,
				SingularTable:       config.NamingStrategy.SingularTable,
				NoLowerCase:         config.NamingStrategy.NoLowerCase,
				IdentifierMaxLength: config.NamingStrategy.IdentifierMaxLength,
			},
			Logger: DefaultLogger,
		}
	}
	return DefaultGORMConfig
}

type IConfig interface {
	GetName() string
	DB() (*gorm.DB, error)
}

type ConfigMysql struct {
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
	Config                        *GORMConfig
}

func (m *ConfigMysql) GetName() string {
	return m.Name
}
func (m *ConfigMysql) DB() (*gorm.DB, error) {
	dsnConf, _ := mysqldriver.ParseDSN(m.DSN)
	return Connect(mysql.New(mysql.Config{
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

type ConfigSQLite struct {
	Name   string
	DSN    string
	Config *GORMConfig
}

func (m *ConfigSQLite) GetName() string {
	return m.Name
}

func (m *ConfigSQLite) DB() (*gorm.DB, error) {
	return Connect(sqlite.Open(m.DSN), m.Config.GORMConfig())
}

type ConfigPostgres struct {
	Name                 string
	DSN                  string
	WithoutQuotingCheck  bool
	PreferSimpleProtocol bool
	WithoutReturning     bool
	Config               *GORMConfig
}

func (m *ConfigPostgres) GetName() string {
	return m.Name
}

func (m *ConfigPostgres) DB() (*gorm.DB, error) {
	return Connect(postgres.New(postgres.Config{
		DSN:                  m.DSN,
		WithoutQuotingCheck:  m.WithoutQuotingCheck,
		PreferSimpleProtocol: m.PreferSimpleProtocol,
		WithoutReturning:     m.WithoutReturning,
	}), m.Config.GORMConfig())
}

type ConfigSQLServer struct {
	Name              string
	DSN               string
	DefaultStringSize int
	Config            *GORMConfig
}

func (m *ConfigSQLServer) GetName() string {
	return m.Name
}

func (m *ConfigSQLServer) DB() (*gorm.DB, error) {
	return Connect(sqlserver.New(sqlserver.Config{
		DSN:               m.DSN,
		DefaultStringSize: m.DefaultStringSize,
	}), m.Config.GORMConfig())
}

type ConfigClickhouse struct {
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
	Config                       *GORMConfig
}

func (m *ConfigClickhouse) GetName() string {
	return m.Name
}

func (m *ConfigClickhouse) DB() (*gorm.DB, error) {
	return Connect(clickhouse.New(clickhouse.Config{
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
