package mysql

import (
	"fmt"
	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/pkg6/gorms"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type NameDBConfig struct {
	User                          string
	Password                      string
	Host                          string
	Port                          int
	Charset                       string
	NameDatabase                  map[string]string
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

func (c *NameDBConfig) NameDB() (map[string]*gorm.DB, error) {
	var (
		err error
	)
	mysqlConfig := mysql.Config{
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
	}
	if c.Port == 0 {
		c.Port = 3306
	}
	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}
	mapDB := map[string]*gorm.DB{}
	for name, database := range c.NameDatabase {
		var db *gorm.DB
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			c.User,
			c.Password,
			c.Host,
			c.Port,
			database,
			c.Charset,
		)
		dsnConf, _ := mysqldriver.ParseDSN(dsn)
		mysqlConfig.DSNConfig = dsnConf
		mysqlConfig.DSN = dsn
		db, err = gorm.Open(mysql.New(mysqlConfig), c.Config.GORMConfig())
		mapDB[name] = db
	}
	return mapDB, err
}
