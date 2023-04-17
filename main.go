package main

import (
	Db "github.com/Juandiaz5673/go-scraper-test-v2/database_wrk"
	Sc "github.com/Juandiaz5673/go-scraper-test-v2/scraper"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Connect to database
	Db.DbConnection()

	// Scraper getting the verse text
	Sc.Scraper()
}
