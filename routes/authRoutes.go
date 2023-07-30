// authRoutes.go

package routes

import (
	authControllers "myapp/controllers/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoginRoutes(e *echo.Echo) {
	e.GET("/login", authControllers.TampilLogin)
}

func LoginUserRoutes(e *echo.Echo) {
	e.POST("/login", authControllers.LoginUser)
}

func TampilRegisRoutes(e *echo.Echo) {
	e.GET("/register", authControllers.TampilRegis)
}

func RegisterRoutes(e *echo.Echo) {
	// Register the validation middleware
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))

	e.POST("/register", authControllers.RegisterUser)
}

func LogoutRoutes(e *echo.Echo) {
	// Register the validation middleware

	e.GET("/logout", authControllers.Logout)
}
