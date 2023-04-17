package database_wrk

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() (*sql.DB, error) {
	Db, err := sql.Open("mysql", "admin:password@tcp(link:3306)/demobibledb")
	if err != nil {
		return nil, err
	}
	return Db, nil
}
