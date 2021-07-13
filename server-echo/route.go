package main

import (
	"encoding/json"
	"go-study/server-echo/apis"
	"go-study/server-echo/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {

	// root
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, wsl Echo World!")
	})

	routingView(e)
	routingApi(e)
}

func routingView(e *echo.Echo) {
	g := e.Group("")

	// user
	{
		g := g.Group("/users")
		g.GET("", views.UsersList)
		// g.GET("", views.GetsUser)
		g.GET("/:id", views.UsersDetail)
		// g.PUT("/:id", views.UpdateUser)
		// g.DELETE("/:id", views.DeleteUser)
	}

	// // route list
	// e.GET("routes", func(c echo.Context) error {
	// 	if data, err := json.MarshalIndent(e.Routes(), "", " "); err == nil {
	// 		return c.String(http.StatusOK, string(data))
	// 	} else {
	// 			return err
	// 	}
	// })
}

func routingApi(e *echo.Echo) {
	g := e.Group("api")

	// user
	{
		g := g.Group("/users")
		g.POST("", apis.SaveUser)
		g.GET("", apis.GetsUser)
		g.GET("/:id", apis.GetUser)
		g.PUT("/:id", apis.UpdateUser)
		g.DELETE("/:id", apis.DeleteUser)
	}

	// images
	{
		g := g.Group("images")
		g.POST("", apis.SaveImage)
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