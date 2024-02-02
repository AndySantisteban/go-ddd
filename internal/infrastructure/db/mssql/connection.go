package mssql

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewConnection(DbName string) (*sql.DB, error) {

	db_name := ""

	if DbName == "" {
		db_name = "Centurion"
	} else {
		db_name = DbName
	}
	dsn := fmt.Sprintf("sqlserver://sysugosrv02:UGO!dev0823@96.88.124.75:1433?database=%s&connection+timeout=30&encrypt=disable", db_name)
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
