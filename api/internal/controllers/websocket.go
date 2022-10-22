package controllers

import (
	"encoding/json"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/entities"

	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketController struct {
	channel chan entities.Candle
}

func NewWebSocketController(channel chan entities.Candle) *WebSocketController {
	return &WebSocketController{
		channel,
	}
}

func (c *WebSocketController) Execute(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("Client Successfully Connected")

	for {
		candle := <-c.channel

		candleJson, _ := json.Marshal(candle)

		err := conn.WriteMessage(websocket.TextMessage, candleJson)

		if err != nil {
			break
		}

	}

}
