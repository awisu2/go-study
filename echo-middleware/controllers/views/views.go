package views

import (
	"go-study/echo-middleware/app"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)


type TemplateData struct {
	Body string
	Data interface{}
	Title string
	Config *app.Config
}

func renderTemplate(c echo.Context, data *TemplateData) error {
	data.Config = app.GetConfig()

	err := c.Render(http.StatusOK, "template", data)
	if err != nil {
		log.Println(err)
	}
	return err
}
