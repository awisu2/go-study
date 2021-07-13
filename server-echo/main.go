package main

import (
	"bytes"
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
	e.Static("/assets", "assets")

	Routing(e)

	seTemplate(e)
	setJson(e)
	setupDB()

	// host情報を削るとwindowsのセキュリティアラートが毎回出る
	e.Logger.Fatal(e.Start("localhost:1323"))
}

// template setting
//
// [template · pkg.go.dev](https://pkg.go.dev/text/template#Must)
//
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
	// template.New: 指定されたnameでユニークに allocates する
	tpl := template.New("base")
	customFuncs := createCustomFuntion(tpl)

	// Render にtemplateをセット
	t := &Template{
		// ここで指定したファイル群がキャッシュされる。
		// template.Must: 生成したテンプレートインスタンスを受け取り、第2引数のerrが != nil のときpnic
		// template.Funcs: 拡張関数の割当 (Parseより前にセットする必要あり)
		// template.ParseGlob: パターンにより対象ファイルからテンプレート情報を生成する
		templates: template.Must(
			tpl.
				Funcs(customFuncs).
				ParseGlob("templates/*.html")),
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

// template用 拡張関数
func createCustomFuntion(t *template.Template) map[string]interface{} {
	// template用 拡張関数
	return map[string]interface{} {
		"hello": func () string {
			return "hello world!"
		},
		// 第一引数にテンプレート名を指定することで、引数からでもtemplateを実行する
		//
		// default の {{template}} では name に変数を割り当てることができず、header, footerを毎回記述する必要がありそうなため追加
		//
		"dynamicTemplate": func(name string, data interface{}) (template.HTML, error) {
				buf := bytes.NewBuffer([]byte{})
				err := t.ExecuteTemplate(buf, name, data)
				if err != nil {
					return "", err
				}
				html := template.HTML(buf.String())
				return html, nil
		},		
	}	
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