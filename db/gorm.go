package gorm

import (
	"fmt"

	"article-service/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbConfig       = config.Config.DB
	mysqlConn      *gorm.DB
	postgresqlConn *gorm.DB
	err            error
	errLog         error
)

// initialize database
func init() {
	setupMysqlConn()
}

// setupMysqlConn: setup mysql database connection using the configuration from config.yml
func setupMysqlConn() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	mysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = mysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}
	mysqlConn.LogMode(true)

	mysqlConn.DB().SetMaxIdleConns(0)
	mysqlConn.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)
}

// MysqlConn: return mysql connection from gorm ORM
func MysqlConn() *gorm.DB {
	return mysqlConn
}
