package config

import (
	"fmt"
	"time"

	"bookstore/constants"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MySqlDB struct {
	*gorm.DB
}

var mysqlDbConn *MySqlDB

// mysqlstructmap is a map of pools
func ConfigureMySql() error {
	// mysqlStructMap := make(map[string]*MySqlDB)

	// TODO: load config from file
	bookStoreMasterDBProps := map[string]string{
		"database": "bookstore",
		"username": "kaushal",
		"password": "password",
		"host":     "localhost",
		"port":     "3306",
	}

	url := getSQLUrl(bookStoreMasterDBProps)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}
	db.DB().SetMaxOpenConns(constants.MaxOpenConnections)
	db.DB().SetMaxIdleConns(constants.MaxIdleConnections)
	db.DB().SetConnMaxLifetime(time.Duration(time.Second * 280))
	mysqlDbConn = &MySqlDB{db}

	return nil
}

func GetMysqlClient() *MySqlDB {
	return mysqlDbConn
}

func getSQLUrl(dbProps map[string]string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s&multiStatements=true",
		dbProps["username"],
		dbProps["password"],
		dbProps["host"],
		dbProps["port"],
		dbProps["database"],
		"Asia%2FCalcutta")
}
