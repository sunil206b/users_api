package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysqlHost   = "MYSQL_DATABASE_HOST"
	mysqlPort   = "MYSQL_DATABASE_PORT"
	mysqlUser   = "MYSQL_DATABASE_USER"
	mysqlPass   = "MYSQL_DATABASE_PASS"
	mysqlSchema = "MYSQL_DATABASE_SCHEMA"
)

func MysqlConn() *sql.DB {
	mysqlHost := os.Getenv(mysqlHost)
	mysqlPort := os.Getenv(mysqlPort)
	mysqlUser := os.Getenv(mysqlUser)
	mysqlPass := os.Getenv(mysqlPass)
	mysqlSchema := os.Getenv(mysqlSchema)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlSchema)
	mysqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = mysqlDB.Ping(); err != nil {
		panic(err)
	}
	return mysqlDB
}
