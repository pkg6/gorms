package gorms

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

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

type INameDBConfig interface {
	NameDB() (map[string]*gorm.DB, error)
}

type IConfig interface {
	GetName() string
	DB() (*gorm.DB, error)
}

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
