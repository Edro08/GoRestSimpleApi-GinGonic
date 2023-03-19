package mysql

import (
	"database/sql"
	"errors"
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

var ErrConnectionDB = errors.New("Servidor de BD no Responde")

func GetmysqlURI() string {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	return mysqlURI
}

func SetupStorage(mysqlURI string) (*sql.DB, error) {

	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return nil, ErrConnectionDB
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}
