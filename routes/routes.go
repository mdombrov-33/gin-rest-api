package routes

import (
	"github.com/gin-gonic/gin"
	middleware "github.com/mdombrov-33/ginrestapi/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// Define the routes for the events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Define the protected routes
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// Define the routes for the users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
