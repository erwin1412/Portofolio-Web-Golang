package controllers

import (
	"context"
	"fmt"
	"html/template"
	"io"
	connection "myapp/database"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Replace this with the base URL of your website

func TampilAddProjectController(c echo.Context) error {
	sess, _ := session.Get("session", c)

	// Check if the user is logged in
	username, loggedIn := sess.Values["username"].(string)
	userID := sess.Values["userID"].(int)
	if !loggedIn {
		// User is not logged in, redirect to the login page
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	tmpl, err := template.ParseFiles("views/myproject.html", "views/nav.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Username": username,
		"userID":   userID,
	}

	// Execute the template with the data
	if err := tmpl.Execute(c.Response(), data); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// AddProjectController adalah fungsi untuk menambahkan proyek baru.

func AddProjectController(c echo.Context) error {
	sess, _ := session.Get("session", c)
	User_id := sess.Values["userID"].(int)
	Title := c.FormValue("title")
	Content := c.FormValue("content")
	StartDate := c.FormValue("start")
	EndDate := c.FormValue("end")
	selectedTechnologies := c.Request().Form["tech[]"]

	startTime, err := time.Parse("2006-01-02", StartDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Format tanggal tidak valid untuk Tanggal Mulai"})
	}

	endTime, err := time.Parse("2006-01-02", EndDate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Format tanggal tidak valid untuk Tanggal Selesai"})
	}

	startDateStr := startTime.Format("2006-01-02")
	endDateStr := endTime.Format("2006-01-02")

	// Dapatkan file yang diunggah
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tidak ada file yang diunggah"})
	}

	// Validate the file extension
	validExtensions := []string{".jpg", ".jpeg", ".png"}
	extension := filepath.Ext(file.Filename)
	isValidExtension := false
	for _, validExt := range validExtensions {
		if extension == validExt {
			isValidExtension = true
			break
		}
	}

	if !isValidExtension {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File format not supported. Only JPG, JPEG, and PNG files are allowed."})
	}

	// Buat nama unik untuk gambar
	timestamp := time.Now().UnixNano()
	uniqueName := fmt.Sprintf("%d%s", timestamp, extension)

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuka file yang diunggah"})
	}
	defer src.Close()

	dstPath := filepath.Join("public", "image", uniqueName)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat file tujuan"})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyalin file"})
	}

	// Connect to the database
	ctx := context.Background()
	conn := connection.Conn

	// Prepare the INSERT statement
	stmt := `INSERT INTO project (title, content, startdate, enddate, imagepath, technologies , user_id) 
             VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Execute the INSERT statement with the project data
	_, err = conn.Exec(ctx, stmt, Title, Content, startDateStr, endDateStr, uniqueName, selectedTechnologies, User_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert project data into the database"})
	}

	// Close the connection explicitly after executing the INSERT statement
	return c.Redirect(http.StatusMovedPermanently, "/project")
}
