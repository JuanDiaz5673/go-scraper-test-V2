package main 

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("GO MySQL Tutorial")


	db, err := sql.Open("mysql", "admin:nacional11@tcp(demobibledb.chctxthfleif.us-east-2.rds.amazonaws.com:3306)/demobibledb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO newkingjamesversion (ID,BookNumber, BookName, ChapterName, ChapterNumber, VerseNumber, VerseText, Testament, BookPreference) VALUES('1', '1', 'Genesis', 'The History of Creation', '1', '1', 'In the beginning God created the heavens and the earth.', 'old', 'New King James Version')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successfully added content to table! ")

}