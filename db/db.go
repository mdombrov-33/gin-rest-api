package db

import (
	"database/sql"

	// go get github.com/mattn/go-sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Open the database
	DB, err = sql.Open("sqlite3", "api.db") // Driver name, path to the database file(will be created if it doesn't exist)

	// Check if the database was opened correctly.
	// Because of Gin, panic will not stop the server from running and will only print the message to the console
	if err != nil {
		panic("Could not connect to the database")
	}

	// Set the maximum number of simultaneously open connections to the database
	DB.SetMaxOpenConns(10)

	// Number of idle connections when the database is not used by anyone
	DB.SetMaxIdleConns(5)

	// Create the tables
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	)`

	// Execute the query
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	// Execute the query
	_, err = DB.Exec(createEventsTable)

	// Check if the table was created correctly
	if err != nil {
		panic("Could not create events table")
		// log.Fatalf("Could not create events table: %v", err)
	}

}
