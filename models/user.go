package models

import "github.com/mdombrov-33/ginrestapi/db"

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

	// Insert values in the same order as the query above
	result, err := sqlStatement.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	// Get the ID of the inserted row, we can do that because we set the ID to AUTOINCREMENT
	userId, err := result.LastInsertId()

	// Set the ID of the user to the ID of the inserted row
	u.ID = userId
	return err
}
