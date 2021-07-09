package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)


func SaveUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetsUser(c echo.Context) error {
	// User ID from path `users/:id`
	log.Println("gdsgasdgasd")
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func UpdateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func DeleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}