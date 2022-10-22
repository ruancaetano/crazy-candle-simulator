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

func HandleWebsocketRequest(channel chan entities.Candle) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println(err)
		}

		log.Println("Client Successfully Connected")

		for {
			candle := <-channel

			candleJson, _ := json.Marshal(candle)

			err := conn.WriteMessage(websocket.TextMessage, candleJson)

			if err != nil {
				break
			}

		}
	}

}
