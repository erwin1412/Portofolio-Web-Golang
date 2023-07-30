package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func ContactRoutes(e *echo.Echo) {
	e.GET("/contact", controllers.TampilContact)
}
