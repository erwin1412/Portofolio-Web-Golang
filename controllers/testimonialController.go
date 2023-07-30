package controllers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func TampilTestimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html", "views/nav.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get the session from the context
	sess, err := session.Get("session", c)
	if err != nil {
		// Session not found, continue rendering the testimonial page without a session
		return tmpl.Execute(c.Response(), nil)
	}

	// Get the username from the session
	username, ok := sess.Values["username"].(string)
	if !ok || username == "" {
		// User is not logged in, continue rendering the testimonial page without a session
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

func GenerateData() []map[string]interface{} {
	datas := []map[string]interface{}{
		{
			"user":   "Erwin",
			"image":  "https://images.unsplash.com/photo-1661956601349-f61c959a8fd4?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1471&q=80",
			"quote":  "Kocak Gaming",
			"rating": 5,
		},
		{
			"user":   "Stefanie",
			"image":  "https://images.unsplash.com/photo-1661956602944-249bcd04b63f?ixlib=rb-4.0.3&ixid=M3wxMjA3fDF8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1470&q=80",
			"quote":  "Kocak Bang",
			"rating": 1,
		},
		{
			"user":   "Quin",
			"image":  "https://images.unsplash.com/photo-1688362378089-23fbdd87eaa3?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1470&q=80",
			"quote":  "Kocak Terbang",
			"rating": 4,
		},
		{
			"user":   "Dumbways",
			"image":  "https://images.unsplash.com/photo-1687902625864-faedb40f83a8?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1473&q=80",
			"quote":  "Kocak Gaming",
			"rating": 5,
		},
	}

	return datas
}

func TampilData(c echo.Context) error {
	datas := GenerateData()
	return c.JSON(http.StatusOK, datas)
}
