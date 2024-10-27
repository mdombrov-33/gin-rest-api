package models

import "time"

type Event struct {
	//`binding:"required"` is a struct tag that tells Gin to check if the field is present in the JSON body
	ID          int
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	//later: Save the event to the database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
