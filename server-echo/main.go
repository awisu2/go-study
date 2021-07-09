package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// static file setting(access, server directory)
	e.Static("/", "assets")

	Routing(e)

	// host情報を削るとwindowsのセキュリティアラートが毎回出る
	e.Logger.Fatal(e.Start("localhost:1323"))
}

