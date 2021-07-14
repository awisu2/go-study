package views

import (
	"log"
	"net/http"

	"go-study/server-echo/models"

	"github.com/labstack/echo/v4"
)

func UsersList(c echo.Context) error {
	log.Println("UsersList")

	var users []*models.User
	// データが無くても配列が返却される
	// Unscopedで削除フラグがたっているレコードも取得
	models.Open().DB.Unscoped().Find(&users)

	return renderTemplate(c, &TemplateData{Body: "users", Data: &struct{
		Users []*models.User
	} {
		users,
	}})
}

func UsersDetail(c echo.Context) error {
	log.Println("UsersDetail")

	// formからの値を取得
	id := c.Param("id")

	if id == "" {
		return renderTemplate(c, &TemplateData{Body: "usersDetail"})
	}

	var user models.User
	models.Open().DB.First(&user, id)

	return renderTemplate(c, &TemplateData{Body: "usersDetail", Data: &struct{
		User *models.User
	} {
		&user,
	}})
}

func UsersSave(c echo.Context) error {
	// formからの値を取得
	id := c.Param("id")
	var user models.User
	if id != "" {
		models.Open().DB.First(&user, id)
	}

	log.Println("UsersCreate")
	return renderTemplate(c, &TemplateData{Body: "usersCreate", Data: struct{
		User models.User
	}{
		user,
	}})
}

func UsersCreatePost(c echo.Context) error {
	log.Println("UsersCreatePost")

	// formからの値を取得
	id := c.FormValue("id")
	userId := c.FormValue("userId")
	name := c.FormValue("name")

	// 更新パラメータの用意
	user := models.User{
		UserId: userId,
		Name: name,
	}

	// save
	db := models.Open()
	if id == "" || id == "0" {
		db.DB.Create(&user)
	} else {
		var _user models.User
		db.DB.First(&_user, id)
		db.DB.Model(&_user).Updates(user)
	}

	// リダイレクト
	err := c.Redirect(http.StatusMovedPermanently, "/users?message=追加しました")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UsersDeletePost(c echo.Context) error {
	// formからの値を取得
	id := c.Param("id")
	if id != "" {
		db := models.Open()
		// struct を指定すると全件削除になるとのこと
		db.DB.Delete(&models.User{}, "Id = ?", id)
	}

	err := c.Redirect(http.StatusMovedPermanently, "/users?message=削除しました")
	if err != nil {
		log.Fatal(err)
	}
	return err
}
