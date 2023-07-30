package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func HomeRoutes(e *echo.Echo) {
	e.GET("/home", controllers.TampilHome)
}
