package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mdombrov-33/ginrestapi/db"
	"github.com/mdombrov-33/ginrestapi/models"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Create a new Gin instance with build-in default middleware and automatic recover from panics
	server := gin.Default()

	// Define the routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event"})
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

	// Dummy ID and UserID
	event.ID = 1
	event.UserID = 1

	// Save the event
	err = event.Save()

	// Check if there was an error saving the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	// 201 status code, send back a message and the created event as JSON
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}
