package main

import (
	"encoding/json"
	"go-study/server-echo/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {

	// root
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, wsl Echo World!")
	})

	// user
	{
		g := e.Group("users")
		g.POST("", controllers.SaveUser)
		g.GET("", controllers.GetsUser)
		g.GET("/:id", controllers.GetUser)
		g.PUT("/:id", controllers.UpdateUser)
		g.DELETE("/:id", controllers.DeleteUser)
	}

	// images
	{
		g := e.Group("images")
		g.POST("", controllers.SaveImage)
	}


	// route list
	e.GET("routes", func(c echo.Context) error {
		if data, err := json.MarshalIndent(e.Routes(), "", " "); err == nil {
			return c.String(http.StatusOK, string(data))
		} else {
				return err
		}
	})
}