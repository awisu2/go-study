package views

import (
	"log"
	"net/http"

	"go-study/server-echo/models"

	"github.com/labstack/echo/v4"
)

func UsersList(c echo.Context) error {
	log.Println("UsersList")
	var users []models.User
	// データが無くても配列が返却される
	models.Open().DB.Find(&users)

	data := struct{
		Users []models.User
	} {
		users,
	}
	return renderTemplate(c, &TemplateData{"users", &data})
}

func UsersDetail(c echo.Context) error {
	log.Println("UsersDetail")
	return renderTemplate(c, &TemplateData{"hello2", nil})
}

func UsersCreate(c echo.Context) error {
	log.Println("UsersCreate")
	return renderTemplate(c, &TemplateData{"usersCreate", nil})
}

func UsersCreatePost(c echo.Context) error {
	log.Println("UsersCreatePost")
	userId := c.FormValue("userId")
	name := c.FormValue("name")

	user := models.User{
		UserId: userId,
		Name: name,
	}

	models.Open().DB.Create(&user)

	err := c.Redirect(http.StatusMovedPermanently, "/users")
	if err != nil {
		log.Fatal(err)
	}
	return err
}