package main

import (
	"go-study/server-echo/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func Routing(e *echo.Echo) {

	// root
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, wsl Echo World!")
	})

	// user
	e.POST("/users", controllers.SaveUser)
	e.GET("/users", controllers.GetsUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	// images
	e.POST("/images", controllers.SaveImage)
}