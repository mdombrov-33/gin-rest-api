package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdombrov-33/ginrestapi/models"
)

func signup(context *gin.Context) {
	var user models.User

	// Bind the JSON body to the user struct
	// Gin will automatically parse the JSON body and bind it to the user struct
	// Client should send a JSON object with the same fields as the user struct
	// If not, Gin will automatically set the missing fields to their zero values
	err := context.ShouldBindJSON(&user)

	// Check if the JSON body was parsed correctly
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// Save the user
	err = user.Save()

	// Check if there was an error saving the user
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	// 201 status code, send back a message and the created user as JSON
	context.JSON(http.StatusCreated, gin.H{"message": "User Created!", "user": user})
}
