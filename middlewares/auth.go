package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdombrov-33/ginrestapi/utils"
)

func Authenticate(context *gin.Context) {
	// Get the token from the Authorization header
	token := context.Request.Header.Get("Authorization")

	// Check if on empty token if client did not send a token
	if token == "" {
		// Abort current request, basically make sure that if something goes wrong we don't continue with the request
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is required"})
	}

	// Check if the token is valid
	userId, err := utils.VerifyToken(token)

	// Abort the request if the token is not valid, basically make sure that if something goes wrong we don't continue with the request
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Token is not valid"})
		return
	}

	// Set the userId to be available in the Gin context(will pass later to the route handlers)
	context.Set("userId", userId)

	// Make sure that next event handler in line will be executed correctly
	context.Next()

}
