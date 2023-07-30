// controllers.go

package controllers

import (
	"context"
	"fmt"
	"html/template"
	connection "myapp/database"
	"myapp/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func TampilLogin(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/login.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil)
}

func LoginUser(c echo.Context) error {
	// Parse form data into LoginForm struct
	loginForm := new(models.LoginForm)
	if err := c.Bind(loginForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the login form data
	if err := c.Validate(loginForm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println("Attempting to authenticate user:", loginForm.Username)

	// Authenticate user using models package's function
	userID, username, err := models.AuthenticateUser(loginForm.Username, loginForm.Password)
	if err != nil {
		// Authentication failed, handle the error and provide appropriate error response
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
	}

	// Authentication succeeded
	fmt.Println("User authenticated:", username)

	// Get the session from the context
	sess, _ := session.Get("session", c)

	// Set the username and user ID in the session
	sess.Values["username"] = username
	sess.Values["userID"] = userID

	// Save the session
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error saving session:", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Redirect to the home page or dashboard
	return c.Redirect(http.StatusFound, "/home")
}

func TampilRegis(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/register.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return tmpl.Execute(c.Response(), nil)
}

// SaveUser saves the user data to the database
func SaveUser(user models.User) error {
	// Hash the user's password before saving it to the database
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	query := `INSERT INTO "user" (username, email, password) VALUES ($1, $2, $3)`
	_, err = connection.Conn.Exec(context.Background(), query, user.Username, user.Email, hashedPassword)
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error saving user to database:", err)
	}
	return err
}

// hashPassword hashes the given password using bcrypt
func hashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// RegisterUser handles the registration process
func RegisterUser(c echo.Context) error {
	// Parse form data into User struct
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the form data
	if err := models.ValidateUser(*user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Check if the password and confirm password match
	if user.Password != user.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, "Password and Confirm Password do not match")
	}

	// Save the user data to the database
	if err := SaveUser(*user); err != nil {
		// Handle the error (e.g., show an error message to the user)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// If registration is successful, redirect to a success page or to the login page
	return c.JSON(http.StatusOK, "Registration successful")
}

// ... (other imports)

func Logout(c echo.Context) error {
	// Get the session from the context
	sess, _ := session.Get("session", c)

	// Clear the session data related to the user
	delete(sess.Values, "username")
	delete(sess.Values, "userID")

	// Save the session to apply the changes
	err := sess.Save(c.Request(), c.Response())
	if err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error saving session:", err)
		return err
	}

	// Redirect to the login page after logout
	return c.Redirect(http.StatusSeeOther, "/login")
}
