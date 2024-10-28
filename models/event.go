package models

import (
	"time"

	"github.com/mdombrov-33/ginrestapi/db"
)

type Event struct {
	//`binding:"required"` is a struct tag that tells Gin to check if the field is present in the JSON body
	ID          int64
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events (title, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)` // protect against SQL injection, actual insert happens below

	// Prepare the query to prevent SQL injection attacks and to improve performance
	sqlStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement connection when the function ends
	defer sqlStatement.Close()

	// Insert values in the same order as the query above
	result, err := sqlStatement.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	// Get the ID of the inserted row, we can do that because we set the ID to AUTOINCREMENT
	id, err := result.LastInsertId()

	e.ID = id

	return err

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	// Can also use db.DB.Exec
	// .Query is used when we just want to get a lot of data back
	// .Exec is used when we have more complex query that changes data in the database
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// Close the rows connection when the function ends
	defer rows.Close()

	var events []Event

	// Loop through the rows
	for rows.Next() {
		var e Event

		// Scan the row and place the values in the struct
		err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		// Append the event to the events slice
		events = append(events, e)
	}

	return events, nil
}
