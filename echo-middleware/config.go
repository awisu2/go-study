package main

import (
	"go-study/echo-middleware/app"
	"go-study/echo-middleware/libs/echo"
)

var echoConfig = echo.Config{
	Address: ":1323",
	Router: router,
	IsCors: true,
	Template: &echo.TemplateConfig{
		Key: "common",
		CreateFuncs: echo.CreateUsableTemplateFuncs,
		GlobPattern: "templates/*.html",
	},
	Static: &echo.StaticConfig{
		Prefix: "/assets",
		Root: "assets",
	},
	IsFixLocationForWindows: true,
}

var appConfig = app.Config{
	Title: "echo sample",
	EchoConfig: &echoConfig,
}
