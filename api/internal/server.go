package internal

import (
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/controllers"
	"githbub.com/ruancaetano/crazy-candle-simulator/api/internal/entities"
	"net/http"
)

type Server struct {
	newCandleChan chan entities.Candle
}

func NewServer(channel chan entities.Candle) *Server {
	server := &Server{
		newCandleChan: channel,
	}

	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	http.HandleFunc("/ws", controllers.HandleWebsocketRequest(s.newCandleChan))
}

func (*Server) Listen(address string) error {
	return http.ListenAndServe(address, nil)
}
