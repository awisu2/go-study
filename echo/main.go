package main

import (
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	setMiddleware(e)
	setRoute(e)

	// windowsの場合、hostをつけないとセキュリティアラートが出る
	uri := ":1323"
	if runtime.GOOS == "windows" {
		uri = "localhost" + uri
	}
	e.Logger.Fatal(e.Start(uri))
}

func setMiddleware(e *echo.Echo) {
	// 定番設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
}

func setRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}