package database_wrk

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Db_connection() {
	Db, err := sql.Open("mysql", "admin:nacional11@tcp(demobibledb.chctxthfleif.us-east-2.rds.amazonaws.com:3306)/demobibledb")
	if err != nil {
		panic(err.Error())
	}
	defer Db.Close()
}
