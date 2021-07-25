package views

import "github.com/labstack/echo/v4"

func Root(c echo.Context) error {
	return renderTemplate(c, &TemplateData{Body: "root"})
}