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

	// Close the statement when the function ends
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

func GetAllEvents() []Event {
	return events
}
