package controllers

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo"
)

func SaveImage(c echo.Context) error {
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
