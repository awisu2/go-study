package views

import "github.com/labstack/echo/v4"

func UsersList(c echo.Context) error {
	return renderTemplate(c, &TemplateData{"hello", nil})
}

func UsersDetail(c echo.Context) error {
	return renderTemplate(c, &TemplateData{"hello2", nil})
}