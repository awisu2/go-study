package views

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	data := map[string]interface{}{"message": "world"}
	return c.Render(http.StatusOK, "index", data)
}