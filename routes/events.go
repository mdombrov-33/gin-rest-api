package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdombrov-33/ginrestapi/models"
)

func getEvents(context *gin.Context) {
	// Get all events
	events, err := models.GetAllEvents()

	// Check if there was an error getting the events
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events"})
		return
	}

	// 200 status code, send back events as JSON(handled automatically by Gin)
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	// Get the event ID from the URL. Name in parentheses should match the parameter(:id) in server.GET
	// Use strconv to convert the string to an integer
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Get the event by ID
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get an event"})
		return
	}

	// 200 status code, send back the event as JSON
	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	// Create a new event
	var event models.Event

	// Bind the JSON body to the event struct
	// Gin will automatically parse the JSON body and bind it to the event struct
	// Client should send a JSON object with the same fields as the event struct
	// If not, Gin will automatically set the missing fields to their zero values
	err := context.ShouldBindJSON(&event)

	// Check if the JSON body was parsed correctly
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// Get the userId from the Gin context
	userId := context.GetInt64("userId")
	event.UserID = userId

	// Save the event
	err = event.Save()

	// Check if there was an error saving the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save an event"})
		return
	}

	// 201 status code, send back a message and the created event as JSON
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}

func updateEvent(context *gin.Context) {
	// Get the event ID from the URL. Name in parentheses should match the parameter(:id) in server.GET
	// Use strconv to convert the string to an integer
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	userId := context.GetInt64("userId")

	// Check if the event exists
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get an event"})
		return
	}

	// User can only update their own events
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to update this event"})
		return
	}

	var updatedEvent models.Event

	updatedEvent.UserID = userId

	// Bind the JSON body to the updatedEvent struct
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// Set the ID of the updated event to the ID from the URL
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update an event"})
		return
	}

	// 200 status code, send back a message and the updated event as JSON
	context.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	// Get the event ID from the URL. Name in parentheses should match the parameter(:id) in server.GET
	// Use strconv to convert the string to an integer
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Get the userId from the Gin context
	userId := context.GetInt64("userId")

	// Get the event by ID
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get an event"})
		return
	}

	// User can only delete their own events
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to delete this event"})
		return
	}

	// Delete the event
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete an event"})
		return
	}

	// 200 status code, send back a message and the deleted event as JSON
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted!", "event": event})
}
