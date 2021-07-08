package main

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, wsl Echo World!")
	})

	// user
	e.POST("/users", saveUser)
	e.GET("/users", getsUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// images
	e.POST("/images", saveImage)

	// host情報を削るとwindowsのセキュリティアラートが毎回出る
	e.Logger.Fatal(e.Start("localhost:1323"))
}

func saveUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getsUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func updateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveImage(c echo.Context) error {
	name := c.FormValue("name")

	image, err:= c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// ファイルの保管
	wd, _ := os.Getwd()
	assestDir := path.Join(wd, "assets")

	err = os.MkdirAll(assestDir, 0755)
	if err != nil {
		return err
	}

	filePath := path.Join(assestDir, name) 
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "image saved!")
}
