package main

import (
	"go-study/echo-middleware/app"
	"go-study/echo-middleware/libs/echo"
)

func main() {
	app.SetConfig(&appConfig)
	echo.StartWithCreate(app.GetConfig().EchoConfig)
}
