package controllers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	connection "myapp/database"
	"myapp/models"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func DisplayProjectController(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/project.html", "views/nav.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	query := `
    SELECT project.id, title, content, startdate, enddate, technologies, imagepath, project.user_id, "user".username
    FROM project
    LEFT JOIN "user" ON project.user_id = "user".id 
    WHERE project.user_id = "user".id
    ORDER BY project.id;
`

	projects, errBlogs := connection.Conn.Query(context.Background(), query)
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
		"Projects": resultProject,
	}

	sess, err := session.Get("session", c)
	if err != nil {
		// Session not found, continue rendering the project page without a session
		err = tmpl.Execute(c.Response(), data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	username, ok := sess.Values["username"].(string)
	if !ok || username == "" {
		// User is not logged in, continue rendering the project page without a session
		err = tmpl.Execute(c.Response(), data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return nil
	}

	userID, ok := sess.Values["userID"].(int)
	if !ok {
		// Handle the error if userID is not found in the session
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	// Now you have the username and userID, you can pass them to the template for rendering
	data["Username"] = username
	data["userID"] = userID

	// Render the project page with the projects data and the username
	err = tmpl.Execute(c.Response(), data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}
