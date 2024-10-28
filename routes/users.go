package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdombrov-33/ginrestapi/models"
	"github.com/mdombrov-33/ginrestapi/utils"
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
	context.JSON(http.StatusCreated, gin.H{"message": "User Created!", "userEmail": user.Email})
}

func login(context *gin.Context) {
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

	// Validate the user credentials
	err = user.ValidateCredentials()

	// Check if there was an error validating the user credentials

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	// 200 status code, send back a JWT token
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
