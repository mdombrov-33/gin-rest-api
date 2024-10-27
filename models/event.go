package models

import "time"

type Event struct {
	ID          int
	Title       string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	// Save the event to the database
	events = append(events, e)
}
