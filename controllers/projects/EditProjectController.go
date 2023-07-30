package controllers

import (
	// Package fmt berisi fungsi dasar untuk formatting I/O (input/output).
	"context"
	"fmt"
	"io"
	"mime/multipart"
	connection "myapp/database"
	"myapp/models"
	"os"
	"strconv"

	// Package template berisi fungsi untuk mengelola template HTML.
	"html/template"
	// Package io menyediakan fungsi dasar untuk melakukan I/O (input/output).

	// Package models berisi definisi model untuk proyek.

	// Package http menyediakan implementasi client dan server untuk protokol HTTP.
	"net/http"
	// Package os menyediakan fungsi-fungsi untuk mengoperasikan sistem operasi, termasuk pembacaan file.

	// Package filepath menyediakan fungsi untuk memanipulasi path file.
	"path/filepath"
	// Package strconv mengimplementasikan konversi antara tipe data string dan tipe data numerik.

	// Package time menyediakan fungsi untuk memanipulasi waktu dan tanggal.
	"time"

	// Package echo adalah web framework berkinerja tinggi yang ringan dan minimalis untuk Go.
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// TampilDetailProjectController adalah fungsi untuk menampilkan halaman detail proyek.

// EditProjectController adalah fungsi untuk menampilkan halaman edit proyek.
func EditProjectController(c echo.Context) error {
	sess, _ := session.Get("session", c)

	// Get the username from the session
	username, loggedIn := sess.Values["username"].(string)

	if !loggedIn {
		// User is not logged in, redirect to the login page
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	userID := sess.Values["userID"].(int)
	tmpl, err := template.New("updateproject.html").Funcs(template.FuncMap{
		"inSlice": inSlice,
	}).ParseFiles("views/updateproject.html", "views/nav.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get the project ID from the route parameter
	id := c.Param("id")

	// Fetch the project data from the database based on the ID
	var project models.Project
	err = connection.Conn.QueryRow(context.Background(),
		"SELECT id, title, content, startdate, enddate, technologies, imagepath FROM project WHERE id=$1", id).
		Scan(&project.ID, &project.Title, &project.Content, &project.StartDate, &project.EndDate, &project.Technologies, &project.ImagePath)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Pass the project data to the template
	dataProject := map[string]interface{}{
		"ID":           project.ID,
		"Title":        project.Title,
		"Content":      project.Content,
		"StartDate":    project.StartDate,
		"EndDate":      project.EndDate,
		"Technologies": project.Technologies,
		"ImagePath":    project.ImagePath,
		"Username":     username, // Add the username to the dataProject map
		"userID":       userID,
	}

	return tmpl.Execute(c.Response(), dataProject)
}

func inSlice(needle string, haystack []string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func generateUniqueName(file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)
	uniqueName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	return uniqueName, nil
}
func UpdateProjectController(c echo.Context) error {
	ID := c.FormValue("ID")
	Title := c.FormValue("Title")
	Content := c.FormValue("Content")
	StartDate := c.FormValue("StartDate")
	EndDate := c.FormValue("EndDate")
	Technologies := c.Request().Form["tech[]"]

	fmt.Println("Received ID:", ID)
	if ID == "" {
		return c.JSON(http.StatusBadRequest, "ID field is empty")
	}

	idToInt, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("gagal conversion ke int")
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	fmt.Println("ID:", idToInt)

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var imagePath string
	if file != nil {
		// Generate a unique file name for the uploaded image
		uniqueName, err := generateUniqueName(file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Save the image to the server
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		dst, err := os.Create("public/image/" + uniqueName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Update the ImagePath field with the new file path
		imagePath = uniqueName
	} else {
		// If no new image is uploaded, retain the existing image path
		existingImagePath, err := getExistingImagePath(idToInt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		imagePath = existingImagePath
	}

	// Rest of the update logic
	// ...

	// Update the project in the database with the new values
	dataUpdate, err := connection.Conn.Exec(context.Background(), "UPDATE project SET title=$1, content=$2, startdate=$3, enddate=$4, technologies=$5, imagepath=$6 WHERE id=$7",
		Title, Content, StartDate, EndDate, Technologies, imagePath, idToInt)
	if err != nil {
		fmt.Println("error guys", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("halo bang", dataUpdate.RowsAffected())

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func getExistingImagePath(projectID int) (string, error) {
	var imagePath string
	err := connection.Conn.QueryRow(context.Background(), "SELECT imagepath FROM project WHERE id=$1", projectID).Scan(&imagePath)
	if err != nil {
		return "", err
	}
	return imagePath, nil
}
