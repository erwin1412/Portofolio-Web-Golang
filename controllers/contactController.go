package controllers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func TampilContact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html", "views/nav.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get the session from the context
	sess, err := session.Get("session", c)
	if err != nil {
		// Session not found, continue rendering the contact page without a session
		return tmpl.Execute(c.Response(), nil)
	}

	// Get the username from the session
	username, ok := sess.Values["username"].(string)
	if !ok || username == "" {
		// User is not logged in, continue rendering the contact page without a session
		return tmpl.Execute(c.Response(), nil)
	}

	// Get the userID from the session
	userID, ok := sess.Values["userID"].(int)
	if !ok {
		// Handle the error if userID is not found in the session
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	data := map[string]interface{}{
		"Username": username,
		"userID":   userID,
	}

	return tmpl.Execute(c.Response(), data)
}
