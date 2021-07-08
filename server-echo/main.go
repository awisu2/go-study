package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	Routing(e)

	// host情報を削るとwindowsのセキュリティアラートが毎回出る
	e.Logger.Fatal(e.Start("localhost:1323"))
}

