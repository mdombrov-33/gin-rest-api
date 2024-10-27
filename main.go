package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdombrov-33/ginrestapi/models"
)

func main() {
	// Create a new Gin instance with build-in default middleware and automatic recover from panics
	server := gin.Default()

	// Define the routes
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	// Get all events
	events := models.GetAllEvents()

	// 200 status code, send back events as JSON(handled automatically by Gin)
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// Create a new event
	var event models.Event

	// Bind the JSON body to the event struct
	// Gin will automatically parse the JSON body and bind it to the event struct
	// Client should send a JSON object with the same fields as the event struct,
	// If not, Gin will automatically set the missing fields to their zero values
	err := context.ShouldBindJSON(&event)

	// Check if the JSON body was parsed correctly
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not pare request data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})

}
