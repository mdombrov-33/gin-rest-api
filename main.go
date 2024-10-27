package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin instance with build-in default middleware and automatic recover from panics
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	//send back data in JSON format
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) // 200 status code, gin.H is a map we can use to send back data
}
