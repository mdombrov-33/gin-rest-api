package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mdombrov-33/ginrestapi/models"
)

func registerForEvent(context *gin.Context) {
	// Get the user ID from the context
	userId := context.GetInt64("userId")

	// Get the event ID from the URL
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Check if event exists
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get an event"})
		return
	}

	err = event.Register(userId)

	// Check if the registration was successful
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event"})
		return
	}

	// 200 status code, send back the message
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for the event"})
}

func cancelRegistrationForEvent(context *gin.Context) {
	// Get the user ID from the context
	userId := context.GetInt64("userId")

	// Get the event ID from the URL
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	// Check if the event ID is valid
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	// Check if the canceling was successful
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}

	// 200 status code, send back the message
	context.JSON(http.StatusOK, gin.H{"message": "Successfully canceled registration"})

}
