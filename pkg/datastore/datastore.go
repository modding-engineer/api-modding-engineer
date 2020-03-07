package datastore

import "database/sql"
import _ "github.com/mattn/go-sqlite3"

func init() {
	database, _ := sql.Open("sqlite3", "./apis.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS apis (id BLOB PRIMARY KEY, name TEXT, url TEXT)")
	statement.Exec()
}
