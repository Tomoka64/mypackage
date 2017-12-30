package main

import (
	"database/sql"

	_ "github.com/github.com/go-sql-driver/mysql"
)

var db *spl.DB
var err error

func main() {
	db, err = sql.Open("mysql", "")
	check(err)
	defer db.Close()

	err = db.Ping()
}
