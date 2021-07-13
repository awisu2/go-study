package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"go-study/server-echo/models"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// static file setting(access, server directory)
	e.Static("/", "assets")

	Routing(e)

	seTemplate(e)
	setJson(e)
	setupDB()

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

// json get/response
type JsonParams struct {
	Id string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func setJson(e *echo.Echo) {
	// try: `http POST http://localhost:1323/json id=123`
	e.POST("/json", func(c echo.Context) error {
		// get json params
		params := JsonParams{}
		if err := c.Bind(&params); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		}
		fmt.Println(params)

		// response json
		return c.JSON(http.StatusOK, params)
	})
}


// setupDB
func setupDB() {
	// httpとは関係ないところでmigrate
	models.AutoMigrate()
}