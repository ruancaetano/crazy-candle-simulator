package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	newCandleChan chan Candle
}

func NewServer(channel chan Candle) *Server {
	server := &Server{
		newCandleChan: channel,
	}

	server.setupRoutes()
	return server
}

func (s *Server) handleWsRequest(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println("Client Successfully Connected")

	for {
		candle := <-s.newCandleChan

		candleJson, _ := json.Marshal(candle)

		err := conn.WriteMessage(websocket.TextMessage, candleJson)

		if err != nil {
			break
		}

	}

}

func (s *Server) setupRoutes() {
	http.HandleFunc("/ws", s.handleWsRequest)
}

func (*Server) Listen(address string) error {
	return http.ListenAndServe(address, nil)
}
