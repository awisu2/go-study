package websocket

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var instanceHub *Hub

func GetHub() *Hub  {
	if instanceHub == nil {
		instanceHub = NewHub()
		go instanceHub.run()
	}
	return instanceHub
}

func ConnectWebSocket(c echo.Context) error {
	// return connectWebSocketGorilla(c)
	return ServeWs(GetHub(), c.Response(), c.Request())
}

// シンプルにgorillaWebsocketに接続
func connectWebSocketGorilla(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer ws.Close()

	for {
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		log.Println(msg)
	}
}

// netに付属のwebsocketでの接続
// import "golang.org/x/net/websocket"
//
// func ConnectWebSocketNet(c echo.Context) {
// 	websocket.Handler(func(ws *websocket.Conn) {
// 		defer CloseWebSocket(ws)
// 		for {
// 			err := websocket.Message.Send(ws, "Hello Client")
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}

// 			msg := ""
// 			err = websocket.Message.Receive(ws, &msg)
// 			if err != nil {
// 				c.Logger().Error(err)
// 			}
// 			log.Printf("%s\n", msg)
// 		}
// 	}).ServeHTTP(c.Response(), c.Request())
// }

// func CloseWebSocket(ws *websocket.Conn) {
// 	log.Println("CloseWebSocket")
// 	ws.Close()
// }