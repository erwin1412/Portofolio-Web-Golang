package controllers

import (
	"context"
	"fmt"
	connection "myapp/database"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteProjectController(c echo.Context) error {
	id := c.Param("id")
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid project ID"})
	}

	Tonn := connection.Conn
	ctx := context.Background()

	// Retrieve the image path from the database for the project
	var imagePath string
	err = Tonn.QueryRow(ctx, "SELECT imagepath FROM project WHERE id=$1", idToInt).Scan(&imagePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve imagepath from the database"})
	}

	// Delete the project from the database
	_, err = Tonn.Exec(ctx, "DELETE FROM project WHERE id=$1", idToInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete project from the database"})
	}

	// Delete the image file from the public/image directory
	err = deleteImageFile(imagePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete image file"})
	}

	return c.Redirect(http.StatusMovedPermanently, "/project")
}

func deleteImageFile(imagePath string) error {
	// Construct the absolute file path to the image
	absolutePath := filepath.Join("public", "image", imagePath)

	// Check if the file exists before attempting to delete
	if _, err := os.Stat(absolutePath); os.IsNotExist(err) {
		return fmt.Errorf("file not found")
	}

	// Delete the image file
	err := os.Remove(absolutePath)
	if err != nil {
		return err
	}

	return nil
}
