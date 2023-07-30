package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func TestimonialRoutes(e *echo.Echo) {
	e.GET("/testimonial", controllers.TampilTestimonial)
}

func DataTestimonial(e *echo.Echo) {
	e.GET("/data", controllers.TampilData)
}
