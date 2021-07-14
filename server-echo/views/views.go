package views

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 一定の引数を強制
type TemplateData struct {
	Body string
	Data interface{}
	Message string
}

func renderTemplate(c echo.Context, data *TemplateData) error {
	if data.Message == "" {
		data.Message = c.FormValue("message")
	}

	err := c.Render(http.StatusOK, "template", data)
	if err != nil {
		log.Println(err)
	}
	return err
}
