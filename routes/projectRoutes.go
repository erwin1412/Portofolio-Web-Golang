package routes

import (
	projectControllers "myapp/controllers/projects"

	"github.com/labstack/echo/v4"
)

func TampilAddProjectRoutes(e *echo.Echo) {
	e.GET("/myproject", projectControllers.TampilAddProjectController)
}

func AddProjectRoutes(e *echo.Echo) {
	e.POST("/add-project", projectControllers.AddProjectController)
}

func DetailProjectRoutes(e *echo.Echo) {
	e.GET("/detail/:id", projectControllers.TampilDetailProjectController)
}

func Project(e *echo.Echo) {
	e.GET("/project", projectControllers.DisplayProjectController)
}

func DeleteProjectRoutes(e *echo.Echo) {
	e.POST("/delete/:id", projectControllers.DeleteProjectController)
}

func EditProjectRoutes(e *echo.Echo) {
	e.GET("/edit/:id", projectControllers.EditProjectController)
}

func UpdateProjectRoutes(e *echo.Echo) {
	e.POST("/update", projectControllers.UpdateProjectController)
}
