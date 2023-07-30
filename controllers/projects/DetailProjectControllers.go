package controllers

import (
	"context"
	"fmt"
	"html/template"
	connection "myapp/database"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func TampilDetailProjectController(c echo.Context) error {
	projectID := c.Param("id")

	tmpl, err := template.ParseFiles("views/detail_project.html", "views/nav.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projectIDInt, err := strconv.Atoi(projectID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid project ID")
	}

	ProjectStruct := models.Project{}
	errQuery := connection.Conn.QueryRow(context.Background(),
		"SELECT id, title, content, startdate, enddate, imagepath, technologies FROM project WHERE id=$1", projectIDInt).
		Scan(&ProjectStruct.ID, &ProjectStruct.Title, &ProjectStruct.Content, &ProjectStruct.StartDate, &ProjectStruct.EndDate, &ProjectStruct.ImagePath, &ProjectStruct.Technologies)

	if errQuery != nil {
		fmt.Println("Error fetching project details:", errQuery)
		return c.JSON(http.StatusNotFound, "Project not found")
	}

	data := map[string]interface{}{
		"Project": ProjectStruct,
	}

	sess, _ := session.Get("session", c)

	// Check if the user is logged in
	userID, ok := sess.Values["userID"].(int)
	if ok {
		// User is logged in, get the username and pass it to the template for rendering
		data["Username"] = sess.Values["username"].(string)
		data["userID"] = userID
	}

	return tmpl.Execute(c.Response(), data)
}
