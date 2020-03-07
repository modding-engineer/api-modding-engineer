package datastore

import (
	"database/sql"
	"flag"
	"fmt"
)
import _ "github.com/mattn/go-sqlite3"

func init() {}

var apiStoreCreateStatement = `create table api_store
      (
      	id BINARY not null
      		constraint api_store_pk
      			primary key,
      	name TEXT not null,
      	url TEXT not null
      );`

var apiStoreCreateIndexStatement = `create unique index api_store_url_uindex
      	on api_store (url);`

func APIS() *sql.DB {
	var defaultPath = "apiDatastore.sqlite"
	if !flag.Parsed() {
		flag.StringVar(&defaultPath, "api-datastore", "apiDatastore.sqlite", "datastore file for APIs")
		flag.Parse()
	}
	database, err := sql.Open("sqlite3", defaultPath)
	// TODO: Gracefully create from internalized SQL
	if err != nil {
		panic(fmt.Errorf("error loading api database: %w", err))
	}
	return database
}
