package main

import (
	"myapp/CORS"
	connection "myapp/database"

	"myapp/routes"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CustomValidator is a custom validator that implements the echo.Validator interface.
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is the method to validate request data using the custom validator.
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// If validation fails, return the error
		return err
	}
	return nil
}

func main() {
	e := echo.New()
	connection.ConectDatabase()
	e.Static("/public", "public")
	// Middleware to enable CORS
	// Add your custom CORS middleware
	e.Use(CORS.HandlerCORS)

	key := []byte("your-secret-key") // Change this to a secure random key
	store := sessions.NewCookieStore(key)
	e.Use(session.Middleware(store))
	e.Validator = &CustomValidator{validator: validator.New()}

	// Add other built-in middleware (e.g., Logger and Recover)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Enable custom error handler
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	// routes
	routes.DataTestimonial(e)
	routes.HomeRoutes(e)
	routes.ContactRoutes(e)
	routes.Project(e)
	routes.EditProjectRoutes(e)
	routes.TampilAddProjectRoutes(e)
	routes.AddProjectRoutes(e)
	routes.DetailProjectRoutes(e)
	routes.TestimonialRoutes(e)
	routes.DeleteProjectRoutes(e)
	routes.UpdateProjectRoutes(e)

	routes.LoginRoutes(e)
	routes.LoginUserRoutes(e)
	routes.TampilRegisRoutes(e)
	routes.RegisterRoutes(e)
	routes.LogoutRoutes(e)

	//port
	e.Logger.Fatal(e.Start("localhost:5000"))
}
