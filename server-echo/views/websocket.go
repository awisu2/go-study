package views

import (
	"go-study/server-echo/dbs"
	"go-study/server-echo/libs/websocket"
	"log"

	"github.com/labstack/echo/v4"
)

func HelloWebsocket(c echo.Context) error {
	return websocket.ConnectWebSocket(c)
}

func WebsocketView(c echo.Context) error {
	log.Println("WebsocketList")

	var users []*dbs.User
	// データが無くても配列が返却される
	// Unscopedで削除フラグがたっているレコードも取得
	dbs.Open("").DB.Unscoped().Find(&users)

	return renderTemplate(c, &TemplateData{Body: "websocketConnect"})
}
