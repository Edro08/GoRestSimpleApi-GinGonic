package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	mysqlURI := GetmysqlURI()
	_, err := SetupStorage(mysqlURI)

	if err != nil {
		t.Errorf(err.Error())
	}
}
