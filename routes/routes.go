package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Define the routes for the events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)       //protected route
	server.PUT("/events/:id", updateEvent)    //protected route
	server.DELETE("/events/:id", deleteEvent) //protected route

	// Define the routes for the users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
