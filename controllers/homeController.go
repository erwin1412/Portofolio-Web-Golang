package controllers

import (
	"context"
	"fmt"
	"html/template"
	connection "myapp/database"
	"myapp/models"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func TampilHome(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html", "views/nav.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	query := `
		SELECT project.id, title, content, startdate, enddate, technologies, imagepath, project.user_id, "user".username
		FROM project
		LEFT JOIN "user" ON project.user_id = "user".id 
		WHERE project.user_id = $1
		ORDER BY project.id;
	`

	// Get the session from the context
	sess, err := session.Get("session", c)
	if err != nil {
		// Handle the error if session retrieval fails
		// For simplicity, let's just display the home page without user data
		return tmpl.Execute(c.Response(), nil)
	}

	// Get the username from the session
	username, ok := sess.Values["username"].(string)
	if !ok || username == "" {
		// User is not logged in, but we will still display the home page without user data
		return tmpl.Execute(c.Response(), nil)
	}

	// Get the userID from the session
	userID, ok := sess.Values["userID"].(int)
	if !ok {
		// Handle the error if userID is not found in the session
		// For simplicity, let's just display the home page without user data
		return tmpl.Execute(c.Response(), nil)
	}

	projects, errBlogs := connection.Conn.Query(context.Background(), query, userID)
	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, errBlogs.Error())
	}

	var resultProject []models.Project
	for projects.Next() {
		var each = models.Project{}

		err := projects.Scan(&each.ID, &each.Title, &each.Content, &each.StartDate, &each.EndDate, &each.Technologies, &each.ImagePath, &each.User_id, &each.Author)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		startTime, err := time.Parse("2006-01-02", each.StartDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format for Start Date"})
		}

		endTime, err := time.Parse("2006-01-02", each.EndDate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format for End Date"})
		}

		duration := endTime.Sub(startTime)
		days := int(duration.Hours() / 24)
		months := int(duration.Hours() / (24 * 30))
		years := int(duration.Hours() / (24 * 365))

		var postDate string
		if days < 1 {
			postDate = "Time has expired"
		} else if days < 30 {
			postDate = fmt.Sprintf("%d days", days)
		} else if days < 60 {
			postDate = "1 month"
		} else if days < 365 {
			postDate = fmt.Sprintf("%d months", months)
		} else {
			postDate = fmt.Sprintf("%d years", years)
		}

		each.PostDate = postDate
		resultProject = append(resultProject, each)
	}

	data := map[string]interface{}{
		"Username": username,
		"userID":   userID,
		"Projects": resultProject,
	}

	return tmpl.Execute(c.Response(), data)
}
