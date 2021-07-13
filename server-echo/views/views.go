package views

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TemplateData struct {
	BodyTemplate string
	Data interface{}
}

func renderTemplate(c echo.Context, data TemplateData) error {
	err := c.Render(http.StatusOK, "template", data)
	if err != nil {
		log.Println(err)
	}
	return err
}
