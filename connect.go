package gorms

import (
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// ConnectClickhouse
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html#Clickhouse
func ConnectClickhouse(dsn string) (*gorm.DB, error) {
	return Connect(clickhouse.Open(dsn), DefaultGORMConfig)
}

// ConnectSQLServer
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html#SQLite
func ConnectSQLServer(dsn string) (*gorm.DB, error) {
	return Connect(sqlserver.Open(dsn), DefaultGORMConfig)
}

// ConnectPostgres
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html#PostgreSQL
func ConnectPostgres(dsn string) (*gorm.DB, error) {
	return Connect(postgres.Open(dsn), DefaultGORMConfig)
}

// ConnectSQLite
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html#SQLite
func ConnectSQLite(dsn string) (*gorm.DB, error) {
	return Connect(sqlite.Open(dsn), DefaultGORMConfig)
}

// ConnectMysql
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html#MySQL
func ConnectMysql(dsn string) (*gorm.DB, error) {
	//https://github.com/go-gorm/mysql
	return Connect(mysql.Open(dsn), DefaultGORMConfig)
}

// Connect
//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func Connect(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(dialector, config)
}
