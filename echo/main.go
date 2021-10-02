package main

import (
	"html/template"
	"io"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go-study/echo/db"
)

func main() {
	e := echo.New()

	setMiddleware(e)
	setTemplate(e)
	setRoute(e)

	_db := db.Open()
	db.Migrate(_db)

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

// templateとしてhtml/templateを使用(echo.Renderer interfaceに合わせた設定)
// [Templates | Echo - High performance, minimalist Go web framework](https://echo.labstack.com/guide/templates/)
type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// テンプレートの登録
func setTemplate(e *echo.Echo) {
	// インタフェースに合わせてRendererを上書き,viewsから呼び出す
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
}

func setRoute(e *echo.Echo) {
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}