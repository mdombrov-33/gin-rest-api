package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mdombrov-33/ginrestapi/db"
	"github.com/mdombrov-33/ginrestapi/routes"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Create a new Gin instance with build-in default middleware and automatic recover from panics
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
