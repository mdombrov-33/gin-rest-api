package models

import (
	"errors"

	"github.com/mdombrov-33/ginrestapi/db"
	"github.com/mdombrov-33/ginrestapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"

	// Prepare the query to prevent SQL injection attacks and to improve performance(prepare is optional, but recommended, can do just .Exec)
	sqlStatement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Close the statement connection when the function ends
	defer sqlStatement.Close()

	// Hash the password before saving it to the database
	hashedPassword, err := utils.HashPassword(u.Password)

	// Check if there was an error hashing the password
	if err != nil {
		return err
	}

	// Insert values in the same order as the query above
	result, err := sqlStatement.Exec(u.Email, hashedPassword)

	// Check if there was an error inserting the values
	if err != nil {
		return err
	}

	// Get the ID of the inserted row, we can do that because we set the ID to AUTOINCREMENT
	userId, err := result.LastInsertId()

	// Set the ID of the user to the ID of the inserted row
	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	// QueryRow returns a single row from the database. Because email is unique, we will guarantee that we will get only one row
	row := db.DB.QueryRow(query, u.Email)

	// Create a variable to store the password from the database
	var retrievedPassword string

	// Scan the row and store the values in the variables
	err := row.Scan(&u.ID, &retrievedPassword)

	// Check if there was an error getting the user from the database
	if err != nil {
		return errors.New("invalid credentials")
	}

	// Compare the password from the database with the password from the user request
	passwordIsValid, err := utils.VerifyPassword(u.Password, retrievedPassword)

	// Check if there was an error comparing the passwords
	if err != nil {
		return err
	}

	// If the password is not valid, return an error
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
