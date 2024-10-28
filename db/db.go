package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error

	// Open the database
	DB, err = sql.Open("sqlite3", "api.db") // Driver name, path to the database file

	// Check if the database was opened correctly.
	// Because of Gin, panic will not stop the server from running and will only print the message to the console
	if err != nil {
		panic("Could not connect to the database")
	}

	// Set the maximum number of simultaneously open connections to the database
	DB.SetMaxOpenConns(10)

	// Number of idle connections when the database is not used by anyone
	DB.SetMaxIdleConns(5)
}
