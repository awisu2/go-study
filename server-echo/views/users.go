package views

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-study/server-echo/dbs"
	"go-study/server-echo/libs"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func UsersList(c echo.Context) error {
	log.Println("UsersList")

	var users []*dbs.User
	// データが無くても配列が返却される
	// Unscopedで削除フラグがたっているレコードも取得
	dbs.Open("").DB.Unscoped().Find(&users)

	return renderTemplate(c, &TemplateData{Body: "users", Data: &struct{
		Users []*dbs.User
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

	var user dbs.User
	dbs.Open("").DB.First(&user, id)

	return renderTemplate(c, &TemplateData{Body: "usersDetail", Data: &struct{
		User *dbs.User
	} {
		&user,
	}})
}

func UsersSave(c echo.Context) error {
	// パスからの値を取得
	id := c.Param("id")
	var user dbs.User
	if id != "" {
		dbs.Open("").DB.First(&user, id)
	}

	log.Println("UsersCreate")
	return renderTemplate(c, &TemplateData{Body: "usersCreate", Data: struct{
		User dbs.User
	}{
		user,
	}})
}

type AnyMap map[string]interface{}

func UsersCreatePost(c echo.Context) error {
	log.Println("UsersCreatePost")

	// formからの値を取得
	data := AnyMap{}
	c.Bind(&data)

	accessedAt, _ := data["accessedAt"].(string)

	t, _ := libs.StrToTime(accessedAt)
	data["accessedAt"] = t
	fmt.Println(t)

	id, idOk := data["id"].(string)

	var user dbs.User
	err := mapstructure.Decode(data, &user)
	if err != nil {
		fmt.Println(err)
		return renderTemplate(c, &TemplateData{Body: "usersCreate", Data: struct{
			User dbs.User
		}{
			user,
		}, })
	}
	fmt.Println(user)

	// save
	db := dbs.Open("")
	if idOk && id != "0" {
		var _user dbs.User
		fmt.Println("!!!update")
		db.DB.First(&_user, id)
		db.DB.Model(&_user).Updates(user)
	} else {
		mapstructure.Decode(data, &user)
		user.AccessedAt = time.Now()
		db.DB.Create(&user)
	}

	// リダイレクト
	err = c.Redirect(http.StatusMovedPermanently, "/users?message=追加しました")
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UsersDeletePost(c echo.Context) error {
	// formからの値を取得
	id := c.Param("id")
	if id != "" {
		db := dbs.Open("")
		// struct を指定すると全件削除になるとのこと
		db.DB.Delete(&dbs.User{}, "Id = ?", id)
	}

	err := c.Redirect(http.StatusMovedPermanently, "/users?message=削除しました")
	if err != nil {
		log.Fatal(err)
	}
	return err
}
