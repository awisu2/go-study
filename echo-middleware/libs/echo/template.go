package echo

import (
	"bytes"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func CreateUsableTemplateFuncs (t *template.Template) map[string]interface{} {
	return map[string]interface{} {
		"dynamicTemplate": CreateDynamicTemplate(t),
	}
}

func CreateDynamicTemplate(t *template.Template) interface{} {
	return func (name string, data interface{}) (template.HTML, error) {
		buf := bytes.NewBuffer([]byte{})
		err := t.ExecuteTemplate(buf, name, data)
		if err != nil {
			return "", err
		}
		html := template.HTML(buf.String())
		return html, nil
	}
}
