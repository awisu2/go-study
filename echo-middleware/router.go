package main

import (
	"go-study/echo-middleware/controllers/views"

	"github.com/labstack/echo/v4"
)

func router(e *echo.Echo) {
	// root
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	e.GET("/", views.Root)
}