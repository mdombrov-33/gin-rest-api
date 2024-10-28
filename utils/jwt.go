package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key used to sign the JWT
const secretKey = "supersecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//don't include password or other sensitive data here
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // token will expire in 2 hours
	})

	// Get string as token and return it
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		//Check if token was signed with the correct secret key
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // returns actual data and a boolean

		// If the signing method is not what we expect, return an error
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	// Check if there was an error parsing the token
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	// Check if the token is valid(signed with the correct secret key)
	tokenIsValid := parsedToken.Valid

	// If the token is not valid, return an error
	if !tokenIsValid {
		return 0, errors.New("token is not valid")
	}

	// Get the token claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// Check if the token claims are valid
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Get the userId from the token claims
	userId := int64(claims["userId"].(float64)) // float64 because that's what the claims are stored as
	return userId, nil                          // convert to int64 manually, if we use int64 instead of float64 we'll get an error when using protected routes
}
