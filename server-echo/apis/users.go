package apis

import (
	"go-study/server-echo/dbs"
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

	// _db := OpenSQLiteConnection()
	// defer _db.Close()

	dbs.TryDB()



	return c.String(http.StatusOK, "hi!")
}

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	var user dbs.User
	// Unscoped(): 削除済みレコードも対象にする
	dbs.Open("").DB.Unscoped().First(&user, id)
	return c.JSON(http.StatusOK, &user)
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