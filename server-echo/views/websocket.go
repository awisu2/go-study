package views

import (
	"go-study/server-echo/libs/websocket"
	"go-study/server-echo/models"
	"log"

	"github.com/labstack/echo/v4"
)

func HelloWebsocket(c echo.Context) error {
	return websocket.ConnectWebSocket(c)
}

func WebsocketView(c echo.Context) error {
	log.Println("WebsocketList")

	var users []*models.User
	// データが無くても配列が返却される
	// Unscopedで削除フラグがたっているレコードも取得
	models.Open().DB.Unscoped().Find(&users)

	return renderTemplate(c, &TemplateData{Body: "websocketConnect"})
}
