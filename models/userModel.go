package models

import (
	"context"
	"fmt"
	connection "myapp/database"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID              int
	Username        string `form:"username" validate:"required"`
	Email           string `form:"email" validate:"required,email"`
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required"`
}

type LoginForm struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// ValidateUser validates the User struct
func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}

// AuthenticateUser performs user authentication
// Replace this implementation with your actual authentication logic
func AuthenticateUser(username, password string) (int, string, error) {
	var dbUsername string
	var dbPasswordHash []byte
	var userID int

	// Query the database to retrieve the user's information
	err := connection.Conn.QueryRow(context.Background(), `SELECT "id", "username", "password" FROM "user" WHERE "username" = $1`, username).Scan(&userID, &dbUsername, &dbPasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, "", fmt.Errorf("authentication failed: user not found")
		}
		return 0, "", fmt.Errorf("authentication failed: %v", err)
	}

	// Compare the provided password with the hashed password from the database
	err = bcrypt.CompareHashAndPassword(dbPasswordHash, []byte(password))
	if err != nil {
		// Passwords do not match, authentication failed
		return 0, "", fmt.Errorf("authentication failed: incorrect password")
	}

	// Authentication succeeded, return the userID and username
	return userID, dbUsername, nil
}
