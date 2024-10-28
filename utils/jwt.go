package utils

import (
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
	return token.SignedString(secretKey)
}
