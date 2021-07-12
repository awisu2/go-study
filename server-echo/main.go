package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// static file setting(access, server directory)
	e.Static("/", "assets")

	Routing(e)

	seTemplate(e)

	// host情報を削るとwindowsのセキュリティアラートが毎回出る
	e.Logger.Fatal(e.Start("localhost:1323"))
}

// template setting
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type User struct {
	Name string
}

func seTemplate(e *echo.Echo) {
	// Render にtemplateをセット
	t := &Template{
		// ここで指定したファイル群がキャッシュされる。
		templates: template.Must(template.ParseGlob("assets/templates/*.html")),
	}
	e.Renderer = t

	// 表示してみる
	e.GET("/template1", func(c echo.Context) error {
		// ファイル名はあまり関係ないファイル内のdefaineで定義される名称が呼び出し対象
		err := c.Render(http.StatusOK, "template01", nil)
		if err != nil {
			log.Println(err)
		}
		return err
	})

	// 他のテンプレート読み込み及び、データの引き渡し
	e.GET("/template2", func(c echo.Context) error {

		jon := User{"jon"}
		data := struct {
			User
		} {
			jon,
		}
		err := c.Render(http.StatusOK, "template02", data)
		if err != nil {
			log.Println(err)
		}
		return err
	})

}