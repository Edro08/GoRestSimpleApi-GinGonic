package storage

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	dbUser = "root"
	dbPass = ""
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "gorestsimpleapi"
)

func GetmysqlURI() string {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	return mysqlURI
}

func SetupStorage(mysqlURI string) (*sql.DB, error) {

	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return nil, sql.ErrConnDone
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
