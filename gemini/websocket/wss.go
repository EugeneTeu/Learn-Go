package websocket

import (
	"log"
	"net/http"

	util "gemini/util"

	"github.com/gorilla/websocket"
)

func OpenWebSocket(symbol string) (c *websocket.Conn) {
	
	c, _, err := websocket.DefaultDialer.Dial(util.GetWebSocketUrl(symbol), nil)
	if err != nil {
		log.Println("something wrong w websocket")
		panic(err.Error())
	}
	return c
}

// websocket
//https://github.com/gorilla/websocket/tree/master/examples/echo
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//... Use conn to send and receive messages.
}

//PUBLIC API REQUEST HEADERS
// GET wss://api.gemini.com/v1/marketdata/BTCUSD
// Connection: Upgrade
// Upgrade: websocket
// Sec-WebSocket-Key: uRovscZjNol/umbTt5uKmw==
// Sec-WebSocket-Version: 13
